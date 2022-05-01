package models

import (
	"time"
)

// TariffType the type for users.tariff_type
type TariffType string

const (
	PrePaid  TariffType = "prepaid"  // PrePaid user is the one who pays upfront for the mobile network services and cannot go beyond the limit
	PostPaid TariffType = "postpaid" // PostPaid user is the one who pays his/her bill in the end of the month.
)

// User the model related to table users
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

// setTimeStamp The function set created_at and updated_at on Insert
func (u *User) setTimeStamp() {
	u.CreatedAt = int(time.Now().UTC().Unix())

	u.UpdatedAt = int(time.Now().UTC().Unix())

	return
}

// updateTime The function set updated_at on Save
func (u *User) updateTime() {
	u.UpdatedAt = int(time.Now().UTC().Unix())
}

// isUpdate The function check if is update on Save
func (u *User) isUpdate() bool {
	if u.Id > 0 {
		return true
	}

	return false
}

// TODO we can add After Update Hook for save all history related to balance OR update balance only paid transaction is success
