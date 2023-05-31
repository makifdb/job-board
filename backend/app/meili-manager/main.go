package main

import (
	"job-site/pkg/database"
	"job-site/pkg/meili"
	"log"
	"os"
	"time"

	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

var Db *gorm.DB
var DbURL = os.Getenv("DB_URL")
var MeiliURL = os.Getenv("MEILI_URL")
var MeiliApiKey = os.Getenv("MEILI_API_KEY")

func main() {
	Db = database.ConnectDB(DbURL)
	M := meili.InitMeili(MeiliURL, MeiliApiKey)

	PushJobsToMeili(Db, M)
}

func PushJobsToMeili(Db *gorm.DB, index *meilisearch.Index) {
	limit := 5
	offset := 0
	for {
		err := meili.LoadJobsToMeili(limit, offset, Db, index)
		if err != nil {
			if err.Error() == "no jobs found" {
				log.Println("no jobs found, offset reset")
				offset = 0
			} else {
				log.Println("error loading jobs to meili :" + err.Error())
				offset += limit
			}
		} else {
			log.Println("jobs loaded to meili")
			offset += limit
		}

		time.Sleep(1 * time.Second)
	}
}
