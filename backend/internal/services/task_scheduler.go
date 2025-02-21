package services

import (
	"log"
	"time"

	"github.com/ISDL-dev/ISDL-Sentinel/backend/internal/repositories"
	"github.com/go-co-op/gocron/v2"
)

func SetLabAssistantScheduler(jst *time.Location) {
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
			log.Printf("Task scheduled for: %v (will execute after %v)：%v\n", scheduledTime, duration, userName)
			time.AfterFunc(duration, func() {
				go NotificationLabAssistantScheduleWithTeams(shiftDate, userName)
			})
		} else {
			log.Printf("The shift date %s is in the past. Skipping...\n", shiftDate)
		}
	}
}

func forceLeavingRoomScheduler() {
	log.Printf("Executing forced exit process...")
	err = repositories.UpdateUserStatusToOutRoom()
	if err != nil {
		log.Fatalf("failed to execute force logout query: %v", err)
		return
	}
	log.Printf("Ending forced exit process")
}

func moveUpGradeScheduler() {
	log.Printf("Updating grades...")
	err = repositories.MoveUpGradeRepository()
	if err != nil {
		log.Fatalf("failed to execute query to move up a grade schedule: %v", err)
		return
	}
	log.Printf("Completed grade updates")
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
			SetLabAssistantScheduler(jst)
		}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = ns.NewJob(
		// (分 時 日 月 曜日)
		gocron.CronJob("59 23 * * *", false),
		gocron.NewTask(func() {
			forceLeavingRoomScheduler()
		}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = ns.NewJob(
		// (分 時 日 月 曜日)
		gocron.CronJob("0 0 1 4 *", false),
		gocron.NewTask(func() {
			moveUpGradeScheduler()
		}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	ns.Start()

	select {}
}
