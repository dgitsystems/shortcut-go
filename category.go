package clubhouse

import (
	"encoding/json"
	"time"
)

type Category struct {
	Archived   bool      `json:"archived"`
	Color      string    `json:"color"`
	CreatedAt  time.Time `json:"created_at"`
	EntityType string    `json:"entity_type"`
	ExternalID int64     `json:"external_id"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (ch *Clubhouse) ListCategories() ([]Category, error) {
	body, err := ch.listResources("categories")
	if err != nil {
		return []Category{}, err
	}
	categories := []Category{}
	json.Unmarshal(body, &categories)
	return categories, nil
}
