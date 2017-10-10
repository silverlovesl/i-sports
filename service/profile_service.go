package service

import "i-sports/model"
import "i-sports/store"

// ProfileService [TODO:Comment]
type ProfileService struct {
	ProfileAccessor *store.ProfileAccessor
}

// NewProfileService [TODO:Comment]
func NewProfileService() *ProfileService {
	return &ProfileService{
		ProfileAccessor: store.NewProfileAccessor(),
	}
}

// DoLogin [TODO:Comment]
func (ps *ProfileService) DoLogin(loginID string, password string) model.Profile {
	entProfile := ps.ProfileAccessor.DoLogin(loginID, password)
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
func (ps *ProfileService) GetProfile(id int) model.Profile {
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
