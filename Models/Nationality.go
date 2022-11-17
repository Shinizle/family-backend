package Models

import "time"

type Nationality struct {
	Id        int64     `gorm: "PrimaryKey" json:"id"`
	Name      string    `gorm: "type:varchar(255); not null;" json:"name"`
	Code      string    `gorm: "type:varchar(255); not null;" json:"code"`
	CreatedAt time.Time `gorm: "type:date" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm: "type:date" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
}
