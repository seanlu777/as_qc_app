package db

import (
	"fmt"
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
