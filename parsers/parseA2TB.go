package parsers

import (
	"as_qc_app/db"
	"as_qc_app/internal/api"
	"as_qc_app/internal/models"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// ParseA2TB parses the A2TB data
func ParseA2TB(data api.PushRecordRequest) ([]db.Test, error) {
	var dbData []db.Test
	for _, d := range data.DataList {
		var a2tb models.A2TB
		// Implement your parsing logic here
		rawData := d.RawData
		a2tb.TagId = strings.ToUpper(rawData[6:12])

		temp, err := strconv.ParseInt(rawData[12:14], 16, 64)
		if err != nil {
			return nil, err
		}
		if temp > 127 {
			temp = temp - 256
		}
		tempDec, err := strconv.ParseInt(rawData[14:16], 16, 64)
		if err != nil {
			return nil, err
		}
		a2tb.Temperature = float64(temp) + float64(tempDec)/100

		pressure, err := strconv.ParseInt(rawData[16:22], 16, 64)
		if err != nil {
			return nil, err
		}
		a2tb.Pressure = int(pressure)

		tagStatusBytes, err := hex.DecodeString(rawData[22:24])
		if err != nil {
			return nil, err
		}
		for _, b := range tagStatusBytes {
			bits := fmt.Sprintf("%08b", b)
			a2tb.CableStatus = bits[0:1] == "1"
			a2tb.TemperatureAlarm = bits[1:2] == "1"
			a2tb.LowBatteryAlarm = bits[2:3] == "1"
		}

		batteryLevel, err := strconv.ParseInt(rawData[24:26], 16, 64)
		if err != nil {
			return nil, err
		}
		a2tb.BatteryLevel = int(batteryLevel)

		counter, err := hex.DecodeString(rawData[26:30])
		if err != nil {
			return nil, err
		}
		for i, j := 0, len(counter)-1; i < j; i, j = i+1, j-1 {
			counter[i], counter[j] = counter[j], counter[i]
		}
		timestamp, err := strconv.ParseInt(hex.EncodeToString(counter), 16, 64)
		if err != nil {
			return nil, err
		}
		a2tb.Timestamp = strconv.FormatInt(timestamp*15, 10)

		a2tb.FirmwareVersion = rawData[32:34]

		dbData = append(dbData, db.Test{
			Station:          data.Station,
			TagId:            a2tb.TagId,
			Temperature:      a2tb.Temperature,
			Pressure:         a2tb.Pressure,
			CableStatus:      a2tb.CableStatus,
			TemperatureAlarm: a2tb.TemperatureAlarm,
			LowBatteryAlarm:  a2tb.LowBatteryAlarm,
			BatteryLevel:     a2tb.BatteryLevel,
			Timestamp:        a2tb.Timestamp,
			FirmwareVersion:  a2tb.FirmwareVersion,
			ReceivedAt:       d.ReceivedAt,
			TenMeterRssi:     d.Rssi,
			TestResult:       d.TestResult,
		})
	}
	return dbData, nil
}
