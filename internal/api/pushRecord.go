package api

import (
	"time"
)

type PushRecordRequest struct {
	Station  string     `json:"station"` // Station name, if not exist, send empty string, not null.
	DataList []DataList `json:"dataList"`
}

type DataList struct {
	RawData    string    `json:"rawData"`
	ReceivedAt time.Time `json:"receivedt"`
	Rssi       int       `json:"rssi"`
	TestResult string    `json:"result"`
}

type PushRecordResponse struct {
	Status  string `json:"status"` // "success" or "error"
	Message string `json:"message"`
}
