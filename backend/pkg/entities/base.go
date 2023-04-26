package entities

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"

	"job-site/pkg/publicid"
)

type Base struct {
	ID         int64          `json:"id" gorm:"primary_key,auto_increment"`
	PublicID   string         `json:"public_id" gorm:"unique_index"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeleteAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Tags       pq.StringArray `gorm:"type:text[]" json:"tags"`
	PublicTags pq.StringArray `gorm:"type:text[]" json:"public_tags"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.PublicID, err = publicid.New()
	return
}
