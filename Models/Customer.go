package Models

import (
	"time"
)

type Customer struct {
	Id            int64        `gorm: "PrimaryKey" json:"id"`
	NationalityID int64        `gorm: "type:int" json:"nationality_id"`
	Name          string       `gorm: "type:varchar(255); not null;" json:"name"`
	Birthday      time.Time    `gorm: "type:date; not null;" json:"birthday"`
	Phone         string       `gorm: "type:varchar(50); not null;" json:"phone"`
	CreatedAt     time.Time    `gorm: "type:date" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time    `gorm: "type:date" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	FamilyList    []FamilyList `gorm: "Foreignkey:CustomerID; association_foreignkey:Id;"`
}
