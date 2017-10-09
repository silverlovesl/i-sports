package service

import "i-sports/model"
import "i-sports/store"

type ProfileService struct {
	ProfileAccessor *store.ProfileAccessor
}

func NewProfileService() *ProfileService {
	return &ProfileService{
		ProfileAccessor: store.NewProfileAccessor(),
	}
}

// Login [TODO:Comment]
func (ps *ProfileService) Login(loginID string, password string) model.Profile {
	entProfile := ps.ProfileAccessor.Login(loginID, password)
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

// Profile [TODO:Comment]
func (ps *ProfileService) Profile(id int) model.Profile {
	entProfile := ps.ProfileAccessor.GetProfile(id)
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
