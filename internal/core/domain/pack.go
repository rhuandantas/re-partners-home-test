package domain

type Pack struct {
	Sizes []int `json:"sizes" xorm:"sizes"`
}
