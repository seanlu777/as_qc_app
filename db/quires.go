package db

import (
	"as_qc_app/internal/api"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func SaveA2TB(data []Test) error {
	// Implement your save logic here

	if DB == nil {
		return fmt.Errorf("DB is nil")
	}

	if len(data) == 0 {
		return fmt.Errorf("no data to save")
	}

	//result := DB.Create(&data)
	//if result.Error != nil {
	//	return fmt.Errorf("failed to save A2TB: %v", result.Error)
	//}

	batchSize := 100
	result := DB.CreateInBatches(&data, batchSize)
	if result.Error != nil {
		return fmt.Errorf("failed to save A2TB: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	fmt.Printf("Successfully saved %d A2TB record\n", result.RowsAffected)

	return nil
}

func GetLatestRecordList(req api.GetLatestRecordListRequest) (api.GetLatestRecordListResponse, error) {
	// Implement your get latest record list logic here
	var resp api.GetLatestRecordListResponse
	var latestDataList []api.LatestDataList
	startAt := req.StartAt
	endAt := req.EndAt
	station := req.Station

	if DB == nil {
		resp = api.GetLatestRecordListResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("DB is nil")
	}

	var testList []Test
	var result *gorm.DB

	if station == "" {
		result = DB.Raw(
			`SELECT DISTINCT ON (tag_id) *
            FROM tests
            WHERE tag_id BETWEEN ? AND ?
            ORDER BY tag_id, received_at DESC
            `, startAt, endAt).Scan(&testList)
	} else {
		result = DB.Raw(
			`SELECT DISTINCT ON (tag_id) *
            FROM tests
            WHERE tag_id BETWEEN ? AND ?
            AND station = ?
            ORDER BY tag_id, received_at DESC
            `, startAt, endAt, station).Scan(&testList)
	}
	if result.Error != nil {
		resp = api.GetLatestRecordListResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("failed to get latest record list: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		resp = api.GetLatestRecordListResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("no rows affected")
	}

	fmt.Println("Successfully retrieved latest record list")

	latestDataList = make([]api.LatestDataList, len(testList))
	for i, test := range testList {
		latestDataList[i] = api.LatestDataList{
			Station:          test.Station,
			TagId:            test.TagId,
			Temperature:      test.Temperature,
			Pressure:         test.Pressure,
			CableStatus:      test.CableStatus,
			TemperatureAlarm: test.TemperatureAlarm,
			LowBatteryAlarm:  test.LowBatteryAlarm,
			BatteryLevel:     test.BatteryLevel,
			Timestamp:        test.Timestamp,
			FirmwareVersion:  test.FirmwareVersion,
			TenMeterRssi:     test.TenMeterRssi,
			TestResult:       test.TestResult,
			ReceivedAt:       test.ReceivedAt,
		}
	}
	resp = api.GetLatestRecordListResponse{
		Status:         "success",
		LatestDataList: latestDataList,
	}
	return resp, nil
}

func GetHistoryData(req api.GetHistoryDataRequest) (api.GetHistoryDataResponse, error) {
	// Implement your get history record list logic here

	var resp api.GetHistoryDataResponse
	var historyList []api.HistoryList

	startAt := req.StartAt
	endAt := req.EndAt
	tagId := strings.ToUpper(req.TagId)
	if DB == nil {
		resp = api.GetHistoryDataResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("DB is nil")
	}

	var testList []Test
	result := DB.Where("received_at BETWEEN ? AND ? AND tag_id = ? ORDER BY received_at DESC", startAt, endAt, tagId).Find(&testList)
	if result.Error != nil {
		resp = api.GetHistoryDataResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("failed to get history record list: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		resp = api.GetHistoryDataResponse{
			Status: "error",
		}
		return resp, fmt.Errorf("no rows affected")
	}
	fmt.Println("Successfully retrieved history record list")

	historyList = make([]api.HistoryList, len(testList))
	for i, test := range testList {
		historyList[i] = api.HistoryList{
			Station:          test.Station,
			TagId:            test.TagId,
			Temperature:      test.Temperature,
			Pressure:         test.Pressure,
			CableStatus:      test.CableStatus,
			TemperatureAlarm: test.TemperatureAlarm,
			LowBatteryAlarm:  test.LowBatteryAlarm,
			BatteryLevel:     test.BatteryLevel,
			Timestamp:        test.Timestamp,
			FirmwareVersion:  test.FirmwareVersion,
			TenMeterRssi:     test.TenMeterRssi,
			TestResult:       test.TestResult,
			ReceivedAt:       test.ReceivedAt,
		}
	}
	resp = api.GetHistoryDataResponse{
		Status:      "success",
		HistoryList: historyList,
	}
	return resp, nil
}
