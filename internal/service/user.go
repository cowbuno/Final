package service

import (
	"forum/models"
)

func (s *service) GetUser(id int) *models.User {
	return nil
}

func (s *service) DeleteSession(token string) error {
	if err := s.repo.DeleteSessionByToken(token); err != nil {
		return err
	}
	return nil
}

func (s *service) Authenticate(email string, password string) (*models.Session, error) {
	userID, err := s.repo.Authenticate(email, password)
	if err != nil {
		return nil, err
	}
	session := models.NewSession(userID)

	if err = s.repo.DeleteSessionByUserID(userID); err != nil {
		return nil, err
	}

	if err = s.repo.CreateSession(session); err != nil {
		return nil, err
	}

	return session, nil
}
func (s *service) CreateUser(user models.User) error {
	err := s.repo.CreateUser(user)
	return err
}
