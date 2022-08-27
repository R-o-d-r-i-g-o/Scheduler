package entity

import (
	"scheduler/common/library"
	"time"

	"gorm.io/gorm"
)

// ---------------------------------< 01º Table >--------------------------------- \\

type User struct {
	Id        uint64            `gorm:"column:id;           primaryKey"`
	Name      string            `gorm:"column:name"`
	Email     string            `gorm:"column:email"`
	Phone     int               `gorm:"column:phone"`
	Password  string            `gorm:"column:password"`
	Indisp    []Indisponibility `gorm:"foreignKey:user_id;  references:id"`
	Us_x_P    []User_x_Pastor   `gorm:"foreignKey:user_id;  references:id"`
	CreatedAt time.Time         `gorm:"column:created_at;   autoCreateTime:true"`
	DeletedAt gorm.DeletedAt    `gorm:"index"`
}

func (User) TableName() string {
	return library.TB_USER
}

// ---------------------------------< 02º Table >--------------------------------- \\

type Indisponibility struct {
	Calendar   time.Time      `gorm:"column:calendar;           primaryKey;    autoCreateTime:false"`
	UserId     uint64         `gorm:"column:user_id;            primaryKey"`
	NotPresent bool           `gorm:"column:not_present"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (Indisponibility) TableName() string {
	return library.TB_INDISPONIBILITY
}

// ---------------------------------< 03º Table >--------------------------------- \\

type User_x_Pastor struct {
	UserId       uint64         `gorm:"column:user_id;            primaryKey"`
	PastorId     uint64         `gorm:"column:pastor_id;          primaryKey"`
	OccupationId uint64         `gorm:"column:occupation_id"`
	CreatedAt    time.Time      `gorm:"column:created_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (User_x_Pastor) TableName() string {
	return library.TB_USER_X_PASTOR
}

// ---------------------------------< 04º Table >--------------------------------- \\

type Occupation struct {
	Id        uint64         `gorm:"column:id;                 primaryKey"`
	Name      string         `gorm:"column:name"`
	Us_x_P    User_x_Pastor  `gorm:"foreignKey:occupation_id;  references:id"`
	CreatedAt time.Time      `gorm:"column:created_at;         autoCreateTime:true"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Occupation) TableName() string {
	return library.TB_OCCUPATION
}

// ---------------------------------< 05º Table >--------------------------------- \\

type Pastor struct {
	Id          uint64          `gorm:"column:id;             primaryKey"`
	Name        string          `gorm:"column:name"`
	Description string          `gorm:"column:description"`
	CreatorId   int             `gorm:"column:creator_id"`
	Us_x_P      []User_x_Pastor `gorm:"foreignKey:pastor_id;  references:id"`
	Month_sch   []MonthSchedule `gorm:"foreignKey:pastor_id;  references:id"`
	P_configs   PastorConfigs   `gorm:"foreignKey:pastor_id;  references:id"`
	CreatedAt   time.Time       `gorm:"column:created_at;     autoCreateTime:true"`
	DeletedAt   gorm.DeletedAt  `gorm:"index"`
}

func (Pastor) TableName() string {
	return library.TB_PASTOR
}

// ---------------------------------< 06º Table >--------------------------------- \\

type PastorConfigs struct {
	PastorId             uint64         `gorm:"column:pastor_id;               primaryKey"`
	OrderByExperience    bool           `gorm:"column:order_by_experience"`
	OrderByGroup         bool           `gorm:"column:order_by_group"`
	ManualConfigInterval bool           `gorm:"column:manual_config_interval"`
	UpdatedAt            time.Time      `gorm:"column:updated_at"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
}

func (PastorConfigs) TableName() string {
	return library.TB_PASTOR_CONFIGS
}

// ---------------------------------< 07º Table >--------------------------------- \\

type MonthSchedule struct {
	Name           string         `gorm:"column:name;               primaryKey"`
	PastorId       uint64         `gorm:"column:pastor_id;          primaryKey"`
	IntervalInDays uint64         `gorm:"column:interval_in_days"`
	Warning        string         `gorm:"column:warning"`
	DailyDuty      []DailyDuty    `gorm:"foreignKey:pastor_id;      references:pastor_id"`
	CreatedAt      time.Time      `gorm:"column:created_at;         autoCreateTime:true"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

func (MonthSchedule) TableName() string {
	return library.TB_MONTH_SCHEDULE
}

// ---------------------------------< 08º Table >--------------------------------- \\

type CelebrationDay struct {
	Calendar     time.Time      `gorm:"column:calendar;           primaryKey;               autoCreateTime:false"`
	WorkDay      bool           `gorm:"column:work_day"`
	Solemnities  string         `gorm:"column:solemnities"`
	LastEditorId uint64         `gorm:"column:last_editor_id"`
	SyncDutyDay  DailyDuty      `gorm:"foreignKey:calendar;       references:calendar"`
	UpdatedAt    time.Time      `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (CelebrationDay) TableName() string {
	return library.TB_CELEBRATION_DAY
}

// ---------------------------------< 09º Table >--------------------------------- \\

type DailyDuty struct {
	Calendar     time.Time      `gorm:"column:calendar;                      primaryKey;    autoCreateTime:false"`
	PastorId     uint64         `gorm:"column:pastor_id;                     primaryKey"`
	NumOfMembers uint64         `gorm:"column:num_of_members"`
	NumOfGroups  uint64         `gorm:"column:num_of_groups"`
	SyncDaySch   DaySchedule    `gorm:"foreignKey:calendar; references:calendar; foreignKey:pastor_id; references:pastor_id"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (DailyDuty) TableName() string {
	return library.TB_DAILY_DUTY
}

// ---------------------------------< 10º Table >--------------------------------- \\

type DaySchedule struct {
	Calendar  time.Time      `gorm:"column:calendar;        primaryKey;    autoCreateTime:false"`
	PastorId  uint64         `gorm:"column:pastor_id;       primaryKey"`
	UserId    uint64         `gorm:"column:user_id"`
	GroupId   uint64         `gorm:"column:group_id"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (DaySchedule) TableName() string {
	return library.TB_DAY_SCHEDULE
}

// ---------------------------------< 11º Table >--------------------------------- \\

type Group struct {
	Id        uint64         `gorm:"column:id;              primaryKey"`
	UserId    uint64         `gorm:"column:user_id"`
	DaySch    []DaySchedule  `gorm:"foreignKey:group_id;    references:id"`
	UpdatedAt time.Time      `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Group) TableName() string {
	return library.TB_GROUP
}

// -----------------------------< General Interface >----------------------------- \\

type Tabler interface {
	TableName() string
}
