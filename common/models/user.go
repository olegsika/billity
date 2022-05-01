package models

import (
	"time"
)

type TariffType string

const (
	PrePaid  TariffType = "prepaid"
	PostPaid TariffType = "postpaid"
)

type User struct {
	TableName  struct{}   `sql:"users" json:"-"`
	Id         int        `sql:"id" json:"id"`
	Name       string     `sql:"name" json:"name"`
	Balance    float64    `sql:"balance" json:"balance"`
	MSISDN     string     `sql:"msisdn" json:"msisdn"`
	TariffType TariffType `sql:"tariff_type" json:"tariff_type"`
	CreatedAt  int        `sql:"created_at" json:"created_at"`
	UpdatedAt  int        `sql:"updated_at" json:"updated_at"`
	DeletedAt  int        `sql:"deleted_at" json:"deleted_at"`
}

func (u *User) setTimeStamp() {
	u.CreatedAt = int(time.Now().UTC().Unix())

	u.UpdatedAt = int(time.Now().UTC().Unix())

	return
}

func (u *User) updateTime() {
	u.UpdatedAt = int(time.Now().UTC().Unix())
}

func (u *User) isUpdate() bool {
	if u.Id > 0 {
		return true
	}

	return false
}

// TODO we can add After Update Hook for save all history related to balance OR update balance only paid transaction is success
