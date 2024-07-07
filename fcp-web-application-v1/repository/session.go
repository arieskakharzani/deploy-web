package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionsRepo struct {
	db *gorm.DB
}

func NewSessionsRepo(db *gorm.DB) *sessionsRepo {
	return &sessionsRepo{db}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}

	return nil
}

func (u *sessionsRepo) DeleteSession(token string) error {
	sessions := model.Session{}
	if result := u.db.Where("token = ?", token).Delete(&sessions); result.Error != nil {
		return result.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	sessions := &model.Session{}
	data := u.db.Debug().Model(sessions).Where("email = ?", session.Email).UpdateColumns(
		map[string]interface{}{
			"token":  session.Token,
			"email":  session.Email,
			"expiry": session.Expiry,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	sessions := model.Session{}
	err := u.db.Where("email = ?", email).Take(&sessions).Error
	if err != nil {
		return model.Session{}, err
	}
	return sessions, nil // TODO: replace this
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	sessions := model.Session{}
	err := u.db.Where("token = ?", token).Take(&sessions).Error
	if err != nil {
		return model.Session{}, err
	}
	return sessions, nil // TODO: replace this
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

func (u *sessionsRepo) GetSessionByEmail(email string) (model.Session, error) {
	var sessions model.Session
	err := u.db.Where("email = ?", email).Take(&sessions).Error
	if err != nil {
		return model.Session{}, err
	}

	return sessions, nil
}
