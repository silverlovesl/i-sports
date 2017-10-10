package store

import (
	"i-sports/entity"
	"i-sports/util"
)

// MeasurementAccessor [Comment]
type measurementAccessor struct {
}

var measurementAccessorIns = new(measurementAccessor)

// GetMeasurementAccessor [Comment]
func GetMeasurementAccessor() *measurementAccessor {
	return measurementAccessorIns
}

// GetMeasurements 測定データ取得
func (ma *measurementAccessor) GetMeasurements(userID int) []entity.Measurement {

	var measurements []entity.Measurement
	engine := util.GetEngine()
	engine.Desc("measure_date").Where("user_id = ?", userID).Find(&measurements)
	return measurements
}
