package scheduler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/robfig/cron"
)

const EVERY_12PM = "0 12 * * *"

/*
	 TODO:
		1. integrating with application class.
		2. using functionality on application class
*/
func StartScheduler() error {
	cronjob := cron.NewWithLocation(time.Local)
	cronjob.AddFunc(EVERY_12PM, func() {
		client := http.DefaultClient
		req, err := http.NewRequest(http.MethodDelete, "http://localhost:5000/v1/transactions/delete/all", nil)
		if err != nil {
			fmt.Println("scheduler error creating request: ", err)
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("scheduler error send request:", err)
		}

		if !statusIsOK(res.StatusCode) {
			fmt.Println("scheduler error response not ok:", err)
		}

		fmt.Println("Success deleting all transactions record")
	})
	cronjob.Start()

	return nil
}

func statusIsOK(statuscode int) bool {
	return statuscode >= 200 && statuscode <= 299
}
