package db

import "time"

type Data struct {
	Id       int    `db:"id" json:"id"`
	Trans_id int    `db:"trans_id" json:"trans_id"`
	Key      string `db:"key" json:"key"`
	Val      string `db:"val" json:"val"`
}

type Trans struct {
	Id     int       `db:"id" json:"id"`
	Entity string    `db:"entity" json:"entity"`
	Desc   string    `db:"descc" json:"desc"`
	Ts     time.Time `db:"ts" json:"ts"`
}
