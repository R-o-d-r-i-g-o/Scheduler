package entity

import (
	"gorm.io/gorm"
)

var occupations = []Occupation{
	{Id: 1, Name: "Leader of pastor(s)"},
	{Id: 2, Name: "experienced member"},
	{Id: 3, Name: "non-experienced member"},
}

func Handler(db *gorm.DB) {
	populateTableOccupation(db)
}

func populateTableOccupation(db *gorm.DB) {
	for _, occupation := range occupations {
		if db.Create(&occupation).Error != nil {
			break
		}
	}
}
