package Models

import "time"

type FamilyList struct {
	Id         int64     `gorm: "PrimaryKey" json:"id"`
	CustomerID int64     `gorm: "type:int" json:"customer_id"`
	Name       string    `gorm: "type:varchar(255); not null;" json:"name"`
	Birthday   time.Time `gorm: "type:date; not null;" json:"birthday"`
	Relation   string    `gorm: "type:varchar(255); not null;" json:"relation"`
	CreatedAt  time.Time `gorm: "type:date" json:"created_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm: "type:date" json:"updated_at" sql:"DEFAULT:CURRENT_TIMESTAMP"`
	Customer   Customer  `gorm: "foreignKey:CustomerID"`
}
