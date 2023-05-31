package meili

import (
	"errors"
	"job-site/pkg/entities"
	"log"

	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

const MEILI_INDEX = "jobs"

func InitMeili(MeiliHost, MeiliApiKey string) *meilisearch.Index {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   MeiliHost,
		APIKey: MeiliApiKey,
	})

	// test connection
	_, err := client.GetVersion()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("meili connected")

	// create index
	_, err = client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        MEILI_INDEX,
		PrimaryKey: "id",
	})
	if err != nil {
		log.Println("Error creating index" + err.Error())
	}

	index := client.Index(MEILI_INDEX)

	return index
}

func UpdateMeili(index *meilisearch.Index, documents []map[string]interface{}) error {

	// add documents to meili
	_, err := index.AddDocuments(documents)
	if err != nil {
		return errors.New("error adding documents to meili " + err.Error())
	}

	return nil
}

func DeleteMeili(index *meilisearch.Index, id string) {
	_, err := index.DeleteDocument(id)
	if err != nil {
		log.Println("Error deleting document from meili " + err.Error())
	}
}

func LoadJobsToMeili(limit, offset int, db *gorm.DB, index *meilisearch.Index) error {

	// get job-board jobs from db
	var jobs []entities.Job
	if err := db.Limit(limit).Offset(offset).Where("source = ?", "job-board").Find(&jobs).Error; err != nil {
		log.Println(err)
		return err
	}

	if len(jobs) == 0 {
		if err := db.Limit(limit).Offset(offset).Find(&jobs).Error; err != nil {
			log.Println(err)
			return err
		}
	}

	// check jobs exist
	if len(jobs) == 0 {
		return errors.New("no jobs found")
	}

	// check meili index exist
	if index == nil {
		return errors.New("meili index not found")
	}

	// check jobs exist in meili
	var meiliJobs []map[string]interface{}
	for _, job := range jobs {
		err := index.GetDocument(job.PublicID, nil, nil)
		if err == nil {
			meiliJobs = append(meiliJobs, map[string]interface{}{
				"id": job.PublicID,
			})
		}
	}

	if len(meiliJobs) == 0 {
		return errors.New("no jobs found in meili")
	}

	// load not exist jobs to meili
	for _, job := range jobs {
		var exist bool
		for _, meiliJob := range meiliJobs {
			if job.PublicID == meiliJob["id"] {
				exist = true
			}
		}
		if !exist {
			document := map[string]interface{}{
				"id":           job.PublicID,
				"title":        job.Title,
				"company_name": job.CompanyName,
				"company_logo": job.CompanyLogo,
				"tags":         job.PublicTags,
				"created_at":   job.CreatedAt,
				"source":       job.Source,
			}

			// add document to meili
			_, err := index.AddDocuments([]map[string]interface{}{document})
			if err != nil {
				log.Println("error adding document to meili " + err.Error())
			}
		}
	}

	return nil
}
