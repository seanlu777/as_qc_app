package db

import (
	"time"
)

// Base model
type BaseModel struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time  `json:"updated_at"`
	//DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

// Test model
type Test struct {
	BaseModel
	Station          string    `json:"station"`
	TagId            string    `json:"tag_id"`
	Temperature      float64   `json:"temperature"`
	Pressure         int       `json:"pressure"`
	CableStatus      bool      `json:"cable_status"`
	TemperatureAlarm bool      `json:"temperature_alarm"`
	LowBatteryAlarm  bool      `json:"low_battery_alarm"`
	BatteryLevel     int       `json:"battery_level"`
	Timestamp        string    `json:"timestamp"`
	FirmwareVersion  string    `json:"version"`
	TenMeterRssi     int       `json:"ten_meter_rssi"`
	TestResult       string    `json:"test_result"`
	ReceivedAt       time.Time `json:"received_at"`
}
