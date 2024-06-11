package entity

import "time"

// TODO: Pervios version
type Order struct {
	ID       string    `json:"id" db:"id"`
	IDPas    string    `json:"id_pas" db:"id_pas"`
	DateTime time.Time `json:"datetime" db:"datetime"`
	Time3    string    `json:"time3" db:"time3"`
	Time4    string    `json:"time4" db:"time4"`
	CatPas   string    `json:"cat_pas" db:"cat_pas"`
	Status   string    `json:"status" db:"status"`
	TPZ      time.Time `json:"tpz" db:"tpz"`
	InspSexM int       `json:"INSP_SEX_M" db:"insp_sex_m,string"`
	InspSexF int       `json:"INSP_SEX_F" db:"insp_sex_f,string"`
	TimeOver string    `json:"TIME_OVER" db:"time_over"`
	IDSt1    string    `json:"id_st1" db:"id_st1"`
	IDSt2    string    `json:"id_st2" db:"id_st2"`
}
