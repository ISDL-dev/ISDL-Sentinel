package services

import (
	"log"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/go-co-op/gocron/v2"
)

func setLabAssistantScheduler(jst *time.Location) {
	log.Printf("Updating lab assistant schedule...")
	now := time.Now()
	month := now.Format("2006-01")

	labAssistantSchedule, err := repositories.GetLabAssistantScheduleRepository(month)
	if err != nil {
		log.Fatalf("failed to execute query to get lab assistant schedule: %v", err)
		return
	}

	for _, schedule := range labAssistantSchedule {
		userName := schedule.UserName
		shiftDate := schedule.ShiftDate

		dateTimeStr := shiftDate + " 10:00"
		layout := "2006-01-02 15:04"
		scheduledTime, err := time.ParseInLocation(layout, dateTimeStr, jst)
		if err != nil {
			log.Fatalf("failed to parse date: %v", err)
		}

		if now.Truncate(time.Minute).Before(scheduledTime.Truncate(time.Minute)) {
			duration := time.Until(scheduledTime)
			log.Printf("Task scheduled for: %v (will execute after %v)\n", scheduledTime, duration)
			time.AfterFunc(duration, func() {
				go NotificationLabAssistantScheduleWithTeams(shiftDate, userName)
			})
		} else {
			log.Printf("The shift date %s is in the past. Skipping...\n", shiftDate)
		}
	}
}

func InitializeTaskScheduler() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		return
	}

	ns, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = ns.NewJob(
		// (分 時 日 月 曜日)
		gocron.CronJob("0 0 1 * *", false),
		gocron.NewTask(func() {
			setLabAssistantScheduler(jst)
		}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	ns.Start()

	select {}
}
