package api

import (
	"time"
)

type GetHistoryDataRequest struct {
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}

type GetHistoryDataResponse struct {
	Status      string        `json:"status"` // "success" or "error"
	HistoryList []HistoryList `json:"historyList"`
}

type HistoryList struct {
	Station          string    `json:"station"`
	TagId            string    `json:"tagId"`
	Temperature      float64   `json:"temperature"`
	Pressure         int       `json:"pressure"`
	CableStatus      bool      `json:"cableStatus"`
	TemperatureAlarm bool      `json:"temperatureAlarm"`
	LowBatteryAlarm  bool      `json:"LowBatteryAlarm"`
	BatteryLevel     int       `json:"batteryLevel"`
	Timestamp        string    `json:"timestamp"`
	FirmwareVersion  string    `json:"FirmwareVersion"`
	TenMeterRssi     int       `json:"tenMeterRssi"`
	Result           bool      `json:"result"`
	ReceivedAt       time.Time `json:"receivedAt"`
}
