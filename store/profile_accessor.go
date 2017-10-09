package store

import (
	"i-sports/entity"
	"i-sports/util"
)

type ProfileAccessor struct {
}

func NewProfileAccessor() *ProfileAccessor {
	return &ProfileAccessor{}
}

// GetProfile [GetProfile]
func (pa *ProfileAccessor) GetProfile(id int) entity.Profile {

	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("id = ?", id).Get(&profile)
	return profile
}

// Login [Comment]
func (pa *ProfileAccessor) Login(loginID string, password string) entity.Profile {
	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("email = ? and password = ?", loginID, password).Find(&profile)
	return profile
}
