package handlers

import (
	"errors"
	"time"

	"github.com/SaroarShahan/event-management/infra/database"
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserID      *int64     `json:"user_id"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (evt *Event) SaveEventsHandler() error {
	if err := database.DB.Create(evt).Error; err != nil {
		return err
	}
	return nil
}

func GetAllEventsHandler() ([]Event, error) {
	var events []Event
	if err := database.DB.Find(&events).Error; err != nil {
		return nil, err
	}
	
	if events == nil {
		events = make([]Event, 0)
	}
	
	return events, nil
}

func GetEventHandler(id int64) (*Event, error) {
	var event Event
	
	if err := database.DB.Take(&event, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("event not found")
		}
		return nil, err
	}

	return &event, nil
}

func UpdateEventHandler(event Event) error {
	if err := database.DB.Model(&Event{}).Where("id = ?", event.ID).Updates(event).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEventHandler(id int64) error {
	if err := database.DB.Delete(&Event{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (event Event) RegisterEventHandler(userID int64) error {
	return CreateRegistration(event.ID, userID)
}

func (event Event) DeleteEventRegistrationHandler(userID int64) error {
	return RemoveRegistration(event.ID, userID)
}
