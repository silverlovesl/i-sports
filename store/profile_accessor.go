package store

import (
	"i-sports/entity"
	"i-sports/util"
)

// ProfileAccessor [TODO:Comment]
type ProfileAccessor struct {
}

// NewProfileAccessor [TODO:Comment]
func NewProfileAccessor() *ProfileAccessor {
	return &ProfileAccessor{}
}

// GetProfile [TODO:Comment]
func (pa *ProfileAccessor) GetProfile(id int) entity.Profile {

	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("id = ?", id).Get(&profile)
	return profile
}

// DoLogin [TODO:Comment]
func (pa *ProfileAccessor) DoLogin(loginID string, password string) entity.Profile {
	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("email = ? and password = ?", loginID, password).Find(&profile)
	return profile
}
