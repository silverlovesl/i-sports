package service

import (
	"i-sports/model"
	"i-sports/store"
)

// MeasurementService [TODO:Comment]
type MeasurementService struct {
	MeasurementAccessor *store.MeasurementAccessor
}

// NewMeasurementService [TODO:Comment]
func NewMeasurementService() *MeasurementService {
	return &MeasurementService{
		MeasurementAccessor: store.NewMeasurementAccessor(),
	}
}

// GetMeasurements [TODO:Comment]
func (mc *MeasurementService) GetMeasurements(userID int) []model.Measurement {

	var measurements []model.Measurement
	entMeasurement := mc.MeasurementAccessor.GetMeasurements(userID)

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
