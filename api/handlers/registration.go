package handlers

import (
	"github.com/SaroarShahan/event-management/infra/database"
	"gorm.io/gorm"
)

type Registration struct {
	gorm.Model
	UserID  int64 `gorm:"not null"`
	EventID uint  `gorm:"not null"`
}

func (registration *Registration) SaveRegistrationHandler() error {
	if err := database.DB.Create(registration).Error; err != nil {
		return err
	}
	return nil
}

func CreateRegistration(eventID uint, userID int64) error {
	registration := Registration{
		UserID:  userID,
		EventID: eventID,
	}
	return registration.SaveRegistrationHandler()
}

func RemoveRegistration(eventID uint, userID int64) error {
	if err := database.DB.Where("event_id = ? AND user_id = ?", eventID, userID).Delete(&Registration{}).Error; err != nil {
		return err
	}
	return nil
}
