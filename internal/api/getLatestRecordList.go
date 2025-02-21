// as_qc_app/internal/api
package api

import (
	"time"
)

type GetLatestRecordListRequest struct {
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
	Station string `json:"station"` // If empty, send empty string, not null. Return all stations.
}

type GetLatestRecordListResponse struct {
	Status         string           `json:"status"` // "success" or "error"
	LatestDataList []LatestDataList `json:"latestDataList"`
}

type LatestDataList struct {
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
