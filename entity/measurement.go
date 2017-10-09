package entity

import "time"

// Measurement [Comment]
type Measurement struct {
	ID             int       `xorm:"'id'"`
	UserID         int       `xorm:"'user_id'"`
	MeasureDate    time.Time `xorm:"'measure_date'"`
	Weight         float64   `xorm:"'weight'"`
	BoneWeight     float64   `xorm:"'bone_weight'"`
	FatWeight      float64   `xorm:"'fat_weight'"`
	Bmi            float64   `xorm:"'bmi'"`
	FatRate        float64   `xorm:"'fat_rate'"`
	WaistHipRate   float64   `xorm:"'waist_hip_rate'"`
	LeftHandMusle  float64   `xorm:"'left_hand_musle'"`
	RightHandMusle float64   `xorm:"'right_hand_musle'"`
	BodyMusle      float64   `xorm:"'body_musle'"`
	LeftLegMusle   float64   `xorm:"'left_leg_musle'"`
	RightLegMusle  float64   `xorm:"'right_leg_musle'"`
	Metabolism     int       `xorm:"'metabolism'"`
	AdjustFat      float64   `xorm:"'adjust_fat'"`
	AdjustWeight   float64   `xorm:"'adjust_weight'"`
	AdjustMusle    float64   `xorm:"'adjust_musle'"`
	Point          float64   `xorm:"'point'"`
	CreateAt       time.Time `xorm:"'create_at'"`
}
