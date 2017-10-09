package model

import "time"

// Measurement [Comment]
type Measurement struct {
	ID             int       `json:"id"`
	UserID         int       `json:"userId"`
	MeasureDate    time.Time `json:"measureDate"`
	Weight         float64   `json:"weight"`
	BoneWeight     float64   `json:"boneWeight"`
	FatWeight      float64   `json:"fatWeight"`
	Bmi            float64   `json:"bmi"`
	FatRate        float64   `json:"fatRate"`
	WaistHipRate   float64   `json:"waistHipRate"`
	LeftHandMusle  float64   `json:"leftHandMusle"`
	RightHandMusle float64   `json:"rightHandMusle"`
	BodyMusle      float64   `json:"bodyMusle"`
	LeftLegMusle   float64   `json:"leftLegMusle"`
	RightLegMusle  float64   `json:"rightLegMusle"`
	Metabolism     int       `json:"metabolism"`
	AdjustFat      float64   `json:"adjustFat"`
	AdjustWeight   float64   `json:"adjustWeight"`
	AdjustMusle    float64   `json:"adjustMusle"`
	Point          float64   `json:"point"`
	CreateAt       time.Time `json:"createAt"`
}
