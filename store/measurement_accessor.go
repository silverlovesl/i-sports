package store

import (
	"i-sports/entity"
	"i-sports/util"
)

// MeasurementAccessor [Comment]
type MeasurementAccessor struct {
}

// NewMeasurementAccessor [Comment]
func NewMeasurementAccessor() *MeasurementAccessor {
	return &MeasurementAccessor{}
}

// GetMeasurements 測定データ取得
func (ma *MeasurementAccessor) GetMeasurements(userID int) []entity.Measurement {

	var measurements []entity.Measurement
	engine := util.GetEngine()
	engine.Desc("measure_date").Where("user_id = ?", userID).Find(&measurements)
	return measurements
}
