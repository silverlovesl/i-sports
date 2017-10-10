package service

import (
	"fmt"
	"i-sports/model"
	"i-sports/store"
	"time"
)

// ProfileService [TODO:Comment]
type profileService struct{}

var profileServiceIns = new(profileService)

// GetProfileServiceIns [TODO:Comment]
func GetProfileServiceIns() *profileService {
	if profileServiceIns == nil {
		time.Sleep(1 * time.Second)
		fmt.Println("A new profileService has been create ")
	}
	return profileServiceIns
}

// DoLogin [TODO:Comment]
func (ps *profileService) DoLogin(loginID string, password string) model.Profile {
	entProfile := store.GetProfileAccessorIns().DoLogin(loginID, password)
	return model.Profile{
		ID:       entProfile.ID,
		Name:     entProfile.Name,
		Gender:   entProfile.Gender,
		Email:    entProfile.Email,
		Height:   entProfile.Height,
		Weight:   entProfile.Weight,
		Birthday: entProfile.Birthday,
	}
}

// GetProfile [TODO:Comment]
func (ps *profileService) GetProfile(id int) model.Profile {
	entProfile := store.GetProfileAccessorIns().GetProfile(id)
	return model.Profile{
		ID:       entProfile.ID,
		Name:     entProfile.Name,
		Gender:   entProfile.Gender,
		Email:    entProfile.Email,
		Height:   entProfile.Height,
		Weight:   entProfile.Weight,
		Birthday: entProfile.Birthday,
	}
}
