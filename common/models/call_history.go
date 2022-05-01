package models

import "time"

type CallHistoryType string

const (
	CallHistoryTypeCall CallHistoryType = "call"
	CallHistoryTypeSms  CallHistoryType = "sms"
)

type CallHistory struct {
	TableName         struct{}        `sql:"call_history" json:"-"`
	Id                int             `sql:"id" json:"id"`
	SourceMsisdn      string          `sql:"source_msisdn" json:"source_msisdn"`
	DestinationMsisdn string          `sql:"destination_msisdn" json:"destination_msisdn"`
	Type              CallHistoryType `sql:"type" json:"type"`
	Duration          int             `sql:"duration" json:"duration"`
	TariffType        TariffType      `sql:"tariff_type" json:"tariff_type"`
	Tariff            float64         `sql:"tariff" json:"tariff"`
	CreatedAt         int             `sql:"created_at" json:"created_at"`
	UpdatedAt         int             `sql:"updated_at" json:"updated_at"`
	DeletedAt         int             `sql:"deleted_at" json:"deleted_at"`
	//
	RequestCost float64 `sql:"request_cost" json:"request_cost"`
	UserBalance float64 `sql:"user_balance" json:"user_balance"`
}

func (u *CallHistory) setTimeStamp() {
	u.CreatedAt = int(time.Now().UTC().Unix())

	u.UpdatedAt = int(time.Now().UTC().Unix())

	return
}

func (u *CallHistory) updateTime() {
	u.UpdatedAt = int(time.Now().UTC().Unix())
}

func (u *CallHistory) isUpdate() bool {
	if u.Id > 0 {
		return true
	}

	return false
}
