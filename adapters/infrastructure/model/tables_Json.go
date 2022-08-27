package model

import (
	"time"
)

type User struct {
	Id       uint64            `json:"id"`
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Phone    int               `json:"phone"`
	Password string            `json:"password"`
	Indisp   []Indisponibility `json:"Indisponibility ,omitempty"`
	Us_x_P   []User_x_Pastor   `json:"User_x_Pastor   ,omitempty"`
}

type Indisponibility struct {
	YearMonthDay time.Time `json:"year_month_day"`
	UserId       uint64    `json:"user_id"`
	NotPresent   bool      `json:"not_present     ,omitempty"`
}

type User_x_Pastor struct {
	UserId       uint64 `json:"user_id"`
	PastorId     uint64 `json:"pastor_id"`
	OccupationId uint64 `json:"occupation_id"`
}

type Occupation struct {
	Id     uint64        `json:"id"`
	Name   string        `json:"name"`
	Us_x_P User_x_Pastor `json:"User_x_Pastor   ,omitempty"`
}

type Pastor struct {
	Id           uint64          `json:"id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	CreatorId    int             `json:"creator_id"`
	CreationDate int             `json:"creation_date"`
	Us_x_P       []User_x_Pastor `json:"User_x_Pastor  ,omitempty"`
	Month_sch    []MonthSchedule `json:"MonthSchedule  ,omitempty"`
	P_configs    PastorConfigs   `json:"PastorConfigs"`
}

type PastorConfigs struct {
	PastorId             uint64 `json:"pastor_id"`
	OrderByExperience    bool   `json:"order_by_experience"`
	OrderByGroup         bool   `json:"order_by_group"`
	ManualConfigInterval bool   `json:"manual_config_interval"`
}

type MonthSchedule struct {
	YearMonth      time.Time   `json:"year_month"`
	PastorId       uint64      `json:"pastor_id"`
	IntervalInDays uint64      `json:"interval_in_days"`
	Warning        string      `json:"warning"`
	DailyDusty     []DailyDuty `json:"DailyDuty          ,omitempty"`
}

type CelebrationDay struct {
	YearMonthDay time.Time `json:"year_month_day"`
	Hour         time.Time `json:"hour"`
	WorkDay      bool      `json:"work_day"`
	Solemnities  string    `json:"solemnities"`
	LastEditorId uint64    `json:"last_editor_id"`
	SyncDutyDay  DailyDuty `json:"DailyDuty           ,omitempty"`
}

type DailyDuty struct {
	YearMonthDay time.Time   `json:"year_month_day"`
	Hour         time.Time   `json:"hour"`
	PastorId     uint64      `json:"pastor_id"`
	NumOfMembers uint64      `json:"num_of_members"`
	NumOfGroups  uint64      `json:"num_of_groups"`
	SyncDaySch   DaySchedule `json:"DaySchedule       ,omitempty"`
}

type DaySchedule struct {
	YearMonthDay time.Time `json:"year_month_day"`
	Hour         time.Time `json:"hour"`
	PastorId     uint64    `json:"pastor_id"`
	UserId       uint64    `json:"user_id"`
	GroupId      uint64    `json:"group_id"`
	MembGroup    []Group   `json:"Group         ,omitempty"`
}

type Group struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
}
