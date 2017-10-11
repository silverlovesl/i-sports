package service

import (
	"i-sports-backend/model"
	"i-sports-backend/store"
)

// MeasurementService [TODO:Comment]
type measurementService struct{}

var measurementServiceIns = new(measurementService)

// GetMeasurementService [TODO:Comment]
func GetMeasurementService() *measurementService {
	return measurementServiceIns
}

// GetMeasurements [TODO:Comment]
func (mc *measurementService) GetMeasurements(userID int) []model.Measurement {

	var measurements []model.Measurement
	entMeasurement := store.GetMeasurementAccessor().GetMeasurements(userID)

	for _, v := range entMeasurement {
		measurements = append(measurements, model.Measurement{
			ID:             v.ID,
			UserID:         v.UserID,
			MeasureDate:    v.MeasureDate,
			Weight:         v.Weight,
			BoneWeight:     v.BoneWeight,
			FatWeight:      v.FatWeight,
			Bmi:            v.Bmi,
			FatRate:        v.FatRate,
			WaistHipRate:   v.WaistHipRate,
			LeftHandMusle:  v.LeftHandMusle,
			RightHandMusle: v.RightHandMusle,
			BodyMusle:      v.BodyMusle,
			LeftLegMusle:   v.LeftLegMusle,
			RightLegMusle:  v.RightLegMusle,
			Metabolism:     v.Metabolism,
			AdjustFat:      v.AdjustFat,
			AdjustWeight:   v.AdjustWeight,
			AdjustMusle:    v.AdjustMusle,
			Point:          v.Point,
			CreateAt:       v.CreateAt,
		})
	}
	return measurements
}
