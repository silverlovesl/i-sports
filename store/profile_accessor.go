package store

import (
	"fmt"
	"i-sports/entity"
	"i-sports/util"
	"time"
)

// ProfileAccessor [TODO:Comment]
type profileAccessor struct{}

// profileAccessorIns [TODO:Comment]
var profileAccessorIns = new(profileAccessor)

// GetProfileAccessorIns [TODO:Comment]
func GetProfileAccessorIns() *profileAccessor {
	if profileAccessorIns == nil {
		time.Sleep(1 * time.Second)
		fmt.Println("A new profileAccessorIns has been create ")
	}
	return profileAccessorIns
}

// GetProfile [TODO:Comment]
func (pa *profileAccessor) GetProfile(id int) entity.Profile {

	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("id = ?", id).Get(&profile)
	return profile
}

// DoLogin [TODO:Comment]
func (pa *profileAccessor) DoLogin(loginID string, password string) entity.Profile {
	var profile entity.Profile
	engine := util.GetEngine()
	engine.Table("profile").Where("email = ? and password = ?", loginID, password).Find(&profile)
	return profile
}
