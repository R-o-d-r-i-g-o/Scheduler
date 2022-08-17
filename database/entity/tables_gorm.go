package entity

import (
	"scheduler/library"
	"time"
)

// ---------------------------------< 01º Table >--------------------------------- \\

type User struct {
	Id       uint64            `gorm:"column:id;           primaryKey"`
	Name     string            `gorm:"column:name"`
	Email    string            `gorm:"column:email"`
	Phone    int               `gorm:"column:phone"`
	Password string            `gorm:"column:password"`
	Indisp   []Indisponibility `gorm:"foreignKey:user_id;  references:id"`
	Us_x_P   []User_x_Pastor   `gorm:"foreignKey:user_id;  references:id"`
}

func (User) TableName() string {
	return library.TB_USER
}

// ---------------------------------< 02º Table >--------------------------------- \\

type Indisponibility struct {
	YearMonthDay time.Time `gorm:"column:year_month_day;     primaryKey;    autoCreateTime:false"`
	UserId       uint64    `gorm:"column:user_id;            primaryKey"`
	NotPresent   bool      `gorm:"column:not_present"`
}

func (Indisponibility) TableName() string {
	return library.TB_INDISPONIBILITY
}

// ---------------------------------< 03º Table >--------------------------------- \\

type User_x_Pastor struct {
	UserId       uint64 `gorm:"column:user_id;            primaryKey"`
	PastorId     uint64 `gorm:"column:pastor_id;          primaryKey"`
	OccupationId uint64 `gorm:"column:occupation_id"`
}

func (User_x_Pastor) TableName() string {
	return library.TB_USER_X_PASTOR
}

// ---------------------------------< 04º Table >--------------------------------- \\

type Occupation struct {
	Id     uint64        `gorm:"column:id;                 primaryKey"`
	Name   string        `gorm:"column:name"`
	Us_x_P User_x_Pastor `gorm:"foreignKey:occupation_id;  references:id"`
}

func (Occupation) TableName() string {
	return library.TB_OCCUPATION
}

// ---------------------------------< 05º Table >--------------------------------- \\

type Pastor struct {
	Id           uint64          `gorm:"column:id;             primaryKey"`
	Name         string          `gorm:"column:name"`
	Description  string          `gorm:"column:description"`
	CreatorId    int             `gorm:"column:creator_id"`
	CreationDate int             `gorm:"column:creation_date"`
	Us_x_P       []User_x_Pastor `gorm:"foreignKey:pastor_id;  references:id"`
	Month_sch    []MonthSchedule `gorm:"foreignKey:pastor_id;  references:id"`
	P_configs    PastorConfigs   `gorm:"foreignKey:pastor_id;  references:id"`
}

func (Pastor) TableName() string {
	return library.TB_PASTOR
}

// ---------------------------------< 06º Table >--------------------------------- \\

type PastorConfigs struct {
	PastorId             uint64 `gorm:"column:pastor_id;               primaryKey"`
	OrderByExperience    bool   `gorm:"column:order_by_experience"`
	OrderByGroup         bool   `gorm:"column:order_by_group"`
	ManualConfigInterval bool   `gorm:"column:manual_config_interval"`
}

func (PastorConfigs) TableName() string {
	return library.TB_PASTOR_CONFIGS
}

// ---------------------------------< 07º Table >--------------------------------- \\

type MonthSchedule struct {
	YearMonth      time.Time   `gorm:"column:year_month;         primaryKey;      autoCreateTime:false"`
	PastorId       uint64      `gorm:"column:pastor_id;          primaryKey"`
	IntervalInDays uint64      `gorm:"column:interval_in_days"`
	Warning        string      `gorm:"column:warning"`
	DailyDusty     []DailyDuty `gorm:"foreignKey:pastor_id;      references:pastor_id"`
}

func (MonthSchedule) TableName() string {
	return library.TB_MONTH_SCHEDULE
}

// ---------------------------------< 08º Table >--------------------------------- \\

type CelebrationDay struct {
	YearMonthDay time.Time `gorm:"column:year_month_day;                      primaryKey;    autoCreateTime:false"`
	Hour         time.Time `gorm:"column:hour;                                primaryKey;    autoCreateTime:false"`
	WorkDay      bool      `gorm:"column:work_day"`
	Solemnities  string    `gorm:"column:solemnities"`
	LastEditorId uint64    `gorm:"column:last_editor_id"`
	SyncDutyDay  DailyDuty `gorm:"foreignKey:year_month_day; references:year_month_day; foreignKey:hour; references:hour"`
}

func (CelebrationDay) TableName() string {
	return library.TB_CELEBRATION_DAY
}

// ---------------------------------< 09º Table >--------------------------------- \\

type DailyDuty struct {
	YearMonthDay time.Time   `gorm:"column:year_month_day;                                       primaryKey;    autoCreateTime:false"`
	Hour         time.Time   `gorm:"column:hour;                                                 primaryKey;    autoCreateTime:false"`
	PastorId     uint64      `gorm:"column:pastor_id;                                            primaryKey"`
	NumOfMembers uint64      `gorm:"column:num_of_members"`
	NumOfGroups  uint64      `gorm:"column:num_of_groups"`
	SyncDaySch   DaySchedule `gorm:"foreignKey:year_month_day; references:year_month_day; foreignKey:hour; references:hour; foreignKey:pastor_id; references:pastor_id"`
}

func (DailyDuty) TableName() string {
	return library.TB_DAILY_DUTY
}

// ---------------------------------< 10º Table >--------------------------------- \\

type DaySchedule struct {
	YearMonthDay time.Time `gorm:"column:year_month_day;  primaryKey;    autoCreateTime:false"`
	Hour         time.Time `gorm:"column:hour;            primaryKey;    autoCreateTime:false"`
	PastorId     uint64    `gorm:"column:pastor_id;       primaryKey"`
	UserId       uint64    `gorm:"column:user_id"`
	GroupId      uint64    `gorm:"column:group_id"`
	MembGroup    []Group   `gorm:"foreignKey:id;          references:group_id"`
}

func (DaySchedule) TableName() string {
	return library.TB_DAY_SCHEDULE
}

// ---------------------------------< 11º Table >--------------------------------- \\

type Group struct {
	Id     uint64 `gorm:"column:id;        primaryKey"`
	UserId uint64 `gorm:"column:user_id"`
}

func (Group) TableName() string {
	return library.TB_GROUP
}

// -----------------------------< General Interface >----------------------------- \\

type Tabler interface {
	TableName() string
}
