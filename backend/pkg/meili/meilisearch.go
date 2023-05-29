package meili

import (
	"encoding/json"
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

func UpdateMeili(index *meilisearch.Index, documents []map[string]interface{}) {

	// print pretty json
	b, err2 := json.MarshalIndent(documents, "", "  ")
	if err2 != nil {
		log.Println("Error marshalling json" + err2.Error())
	}
	log.Println(string(b))

	_, err := index.AddDocuments(documents)
	if err != nil {
		log.Println("Error adding documents to meili" + err.Error())
	}
}

func DeleteMeili(index *meilisearch.Index, id string) {
	_, err := index.DeleteDocument(id)
	if err != nil {
		log.Println("Error deleting document from meili" + err.Error())
	}
}

func LoadJobsToMeili(limit, offset int, db *gorm.DB, index *meilisearch.Index) error {

	// get jobs from db
	var jobs []entities.Job
	if err := db.Limit(limit).Offset(offset).Find(&jobs).Error; err != nil {
		log.Println(err)
		return err
	}

	// load jobs to meili
	var documents []map[string]interface{}
	for _, job := range jobs {
		documents = append(documents, map[string]interface{}{
			"id":           job.PublicID,
			"title":        job.Title,
			"company_name": job.CompanyName,
			"company_logo": job.CompanyLogo,
			"tags":         job.PublicTags,
			"created_at":   job.CreatedAt,
			"source":       job.Source,
		})
	}

	UpdateMeili(index, documents)

	return nil
}
