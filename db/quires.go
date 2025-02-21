package db

import (
	"as_qc_app/internal/api"
	"fmt"

	"gorm.io/gorm"
)

func SaveA2TB(data []Test) error {
	// Implement your save logic here

	if DB == nil {
		return fmt.Errorf("DB is nil")
	}

	result := DB.Create(&data)
	if result.Error != nil {
		return fmt.Errorf("failed to save A2TB: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	fmt.Println("Successfully saved A2TB record:", data)

	return nil
}

func GetLatestRecordList(req api.GetLatestRecordListRequest) ([]Test, error) {
	// Implement your get latest record list logic here
	startAt := req.StartAt
	endAt := req.EndAt
	station := req.Station

	if DB == nil {
		return nil, fmt.Errorf("DB is nil")
	}

	var testList []Test
	var result *gorm.DB

	if station == "" {
		result = DB.Raw(
			`SELECT DISTINCT ON (tag_id) *
            FROM tests
            WHERE received_at BETWEEN ? AND ?
            ORDER BY tag_id, received_at DESC
            `, startAt, endAt).Scan(&testList)
	} else {
		result = DB.Raw(
			`SELECT DISTINCT ON (tag_id) *
            FROM tests
            WHERE received_at BETWEEN ? AND ?
            AND station = ?
            ORDER BY tag_id, received_at DESC
            `, station, startAt, endAt).Scan(&testList)
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get latest record list: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no rows affected")
	}
	fmt.Println("Successfully retrieved latest record list:", testList)

	return testList, nil
}
