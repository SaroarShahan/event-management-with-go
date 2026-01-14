package handlers

import (
	"time"

	"github.com/SaroarShahan/event-management/infra/database"
)

type Event struct {
	ID 	int64    `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Location string `json:"location"`
	Datetime time.Time `json:"datetime"`
	UserID int64 `json:"user_id"`
}

func (evt Event) SaveEventsHandler() error {
	query := `INSERT INTO events (name, description, location, datetime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(evt.Name, evt.Description, evt.Location, evt.Datetime, evt.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	evt.ID = id

	return err
}

func GetAllEventsHandler() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := database.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := make([]Event, 0)

	for rows.Next() {
		var evt Event
		err := rows.Scan(&evt.ID, &evt.Name, &evt.Description, &evt.Location, &evt.Datetime, &evt.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, evt)
	}

	return events, nil

}

func GetEventHandler(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := database.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Datetime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func UpdateEventHandler(event Event) error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, datetime = ? WHERE id = ?`
	stmt, err := database.DB.Prepare(query)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.Datetime, event.ID)

	return err
}

func DeleteEventHandler(id int64) error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}