package infrastructures

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	model "github.com/ISDL-dev/ISDL-Sentinel/backend/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func osUserCacheDir() string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Caches")
	case "linux", "freebsd":
		return filepath.Join(os.Getenv("HOME"), ".cache")
	}
	log.Printf("TODO: osUserCacheDir on GOOS %q", runtime.GOOS)
	return "."
}

func tokenCacheFile(config *oauth2.Config) string {
	hash := fnv.New32a()
	hash.Write([]byte(config.ClientID))
	hash.Write([]byte(config.ClientSecret))
	hash.Write([]byte(strings.Join(config.Scopes, " ")))
	fn := fmt.Sprintf("go-api-demo-tok%v", hash.Sum32())
	return filepath.Join(osUserCacheDir(), url.QueryEscape(fn))
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := new(oauth2.Token)
	err = gob.NewDecoder(f).Decode(t)
	return t, err
}

func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		log.Printf("Warning: failed to cache oauth token: %v", err)
		return
	}
	defer f.Close()
	gob.NewEncoder(f).Encode(token)
}

func newOAuthClient(ctx context.Context, config *oauth2.Config) *http.Client {
	data, err := loadToken("internal/infrastructures/credentials/token.json")
	if err != nil {
		log.Fatalf("Failed to load token: %v", err)
	}

	var tokenData model.Token
	if err := json.Unmarshal(data, &tokenData); err != nil {
		log.Fatalf("Failed to unmarshal token data: %v", err)
	}

	// oauth2.Token オブジェクトに変換する
	token := &oauth2.Token{
		AccessToken:  tokenData.AccessToken,
		TokenType:    tokenData.TokenType,
		RefreshToken: tokenData.RefreshToken,
		Expiry:       parseExpiry(tokenData.Expiry),
	}

	// トークンが期限切れならリフレッシュする
	if token.Expiry.Before(time.Now()) {
		fmt.Println("token expired, refreshing...")
		token = refreshToken(token, config)
		saveToken("token.json", token)
	}

	return config.Client(ctx, token)
}

func parseExpiry(expiryStr string) time.Time {
	expiry, err := time.Parse(time.RFC3339, expiryStr)
	if err != nil {
		log.Fatalf("Failed to parse expiry time: %v", err)
	}
	return expiry
}

func tokenFromWeb(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	ch := make(chan string)
	randState := fmt.Sprintf("st%d", time.Now().UnixNano())
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/favicon.ico" {
			http.Error(rw, "", 404)
			return
		}
		if req.FormValue("state") != randState {
			log.Printf("State doesn't match: req = %#v", req)
			http.Error(rw, "", 500)
			return
		}
		if code := req.FormValue("code"); code != "" {
			fmt.Fprintf(rw, "<h1>Success</h1>Authorized.")
			rw.(http.Flusher).Flush()
			ch <- code
			return
		}
		log.Printf("no code")
		http.Error(rw, "", 500)
	}))
	defer ts.Close()

	config.RedirectURL = ts.URL
	authURL := config.AuthCodeURL(randState)
	go openURL(authURL)
	log.Printf("Authorize this app at: %s", authURL)
	code := <-ch
	log.Printf("Got code: %s", code)

	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Token exchange error: %v", err)
	}
	return token
}

func openURL(url string) {
	try := []string{"xdg-open", "google-chrome", "open"}
	for _, bin := range try {
		err := exec.Command(bin, url).Run()
		if err == nil {
			return
		}
	}
	log.Printf("Error opening URL in browser.")
}

func loadToken(path string) (*oauth2.Token, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tok oauth2.Token
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tok)
	if err != nil {
		return nil, err
	}

	return &tok, nil
}

func refreshToken(tok *oauth2.Token, config *oauth2.Config) *oauth2.Token {
	tokenSource := config.TokenSource(context.Background(), tok)
	newTok, err := tokenSource.Token()
	if err != nil {
		log.Fatalf("Failed to refresh token: %v", err)
	}
	return newTok
}

func valueOrFileContents(value string, filename string) string {
	if value != "" {
		return value
	}
	slurp, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading %q: %v", filename, err)
	}
	return strings.TrimSpace(string(slurp))
}

func GetCalendarList() {
	calendarIDs := []model.CalendarRoomId{
		model.CalendarRoomId{RoomName: "KC101-large", CalendarId: model.KC101_LARGE_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC103", CalendarId: model.KC103_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC111", CalendarId: model.KC111_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC116", CalendarId: model.KC116_CALENDAR_ID},
		model.CalendarRoomId{RoomName: "KC119", CalendarId: model.KC119_CALENDAR_ID},
	}

	ctx := context.Background()
	b, err := os.ReadFile("internal/infrastructures/credentials/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	c := newOAuthClient(ctx, config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(c))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	for _, room := range calendarIDs {
		events, err := srv.Events.List(room.CalendarId).ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		}
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found.")
		} else {
			for _, item := range events.Items {
				date := item.Start.DateTime
				if date == "" {
					date = item.Start.Date
				}
				fmt.Printf("%v (%v)\n", item.Summary, date)
				fmt.Printf("使用者\n")
				for _, attendee := range item.Attendees {
					fmt.Printf("(%s)\n", attendee.Email)
				}
			}
		}
	}
}
