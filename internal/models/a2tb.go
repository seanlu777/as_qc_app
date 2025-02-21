// internal/api/a2tb.go
package models

type A2TB struct {
	TagId            string  `json:"tagId"`
	Temperature      float64 `json:"temperature"`
	Pressure         int     `json:"pressure"`
	CableStatus      bool    `json:"cableStatus"`
	TemperatureAlarm bool    `json:"temperatureAlarm"`
	LowBatteryAlarm  bool    `json:"low_batteryAlarm"`
	BatteryLevel     int     `json:"batteryLevel"`
	Timestamp        string  `json:"timestamp"`
	FirmwareVersion  string  `json:"version"`
}
