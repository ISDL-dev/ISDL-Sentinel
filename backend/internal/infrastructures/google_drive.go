package infrastructures

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func getClient(config *oauth2.Config) *http.Client {
	tokFile := "internal/infrastructures/credentials/google_drive_token.json"
	tok, err := tokenFromFileForDrive(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveTokenForDrive(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func tokenFromFileForDrive(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func saveTokenForDrive(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

func UploadAvatarFile(avatarFile *multipart.FileHeader) (string, error) {
	file, err := avatarFile.Open()
	if err != nil {
		return "", fmt.Errorf("Failed to open the file: %w", err)
	}
	defer file.Close()

	ctx := context.Background()
	b, err := os.ReadFile("internal/infrastructures/credentials/google_drive_credentials.json")
	if err != nil {
		return "", fmt.Errorf("Failed to read the client secret file: %w", err)
	}

	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return "", fmt.Errorf("Failed to parse the client secret file: %w", err)
	}

	client := getClient(config)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return "", fmt.Errorf("Failed to initialize Google Drive client: %w", err)
	}

	fileMetadata := &drive.File{
		Name:    avatarFile.Filename,
		Parents: []string{os.Getenv("GOOGLE_DRIVE_FOLDER_ID")},
	}

	driveFile, err := srv.Files.Create(fileMetadata).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("Failed to upload the file to Google Drive: %w", err)
	}

	avatarImgPath := fmt.Sprintf("https://drive.google.com/uc?id=%s", driveFile.Id)

	return avatarImgPath, nil
}
