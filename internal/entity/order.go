package entity

import "time"

// TODO: Pervios version
type Order struct {
	ID       int       `json:"id" db:"id"`
	IDPas    string    `json:"id_pas" db:"id_pas"`
	DateTime time.Time `json:"datetime" db:"datetime"`
	Time3    string    `json:"time3" db:"time3"`
	Time4    string    `json:"time4" db:"time4"`
	CatPas   string    `json:"cat_pas" db:"cat_pas"`
	StatusID int       `json:"status_id" db:"status_id"`
	TPZ      time.Time `json:"tpz" db:"tpz"`
	InspSexM int       `json:"insp_sex_m" db:"insp_sex_m"`
	InspSexF int       `json:"insp_sex_f" db:"insp_sex_f"`
	TimeOver string    `json:"time_over" db:"time_over"`
	IDSt1    string    `json:"id_st1" db:"id_st1"`
	IDSt2    string    `json:"id_st2" db:"id_st2"`
}

type Status struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
