package infrastructure

import (
	"fmt"
	"log"
	model "scheduler/adapters/infrastructure/entity"
	"scheduler/common/configs"

	"gorm.io/gorm"
)

var (
	db *gorm.DB

	tables = func(db *gorm.DB) error {
		return db.AutoMigrate(

			// THERE'RE NOT external dependences (foreign keys)
			&model.User{},
			&model.Indisponibility{},
			&model.Pastor{},
			&model.PastorConfigs{},
			&model.Occupation{},
			&model.User_x_Pastor{},
			&model.CelebrationDay{},
			&model.MonthSchedule{},
			&model.DailyDuty{},
			&model.Group{},
			&model.DaySchedule{},

			// THERE'RE external dependences (foreign keys)
		)
	}
)

func GetGormDB() *gorm.DB {
	return db
}

func StartDatabase() {
	databaseConfiguration := createDatabaseStringConfig()
	mysql := NewMysql(databaseConfiguration)

	if err := mysql.connect(); err != nil {
		log.Fatal(err)
	}

	if tables(mysql.db) == nil {
		db = mysql.db
	}
}

func createDatabaseStringConfig() string {
	databaseStringConfig := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=%s&loc=%s",

		configs.DBase.USER,
		configs.DBase.PASSWORD,
		configs.DBase.HOST,
		configs.DBase.PORT,
		configs.DBase.DB,

		configs.Time.PARSE,
		configs.Time.LOCATION,
	)

	return databaseStringConfig
}
