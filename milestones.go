package clubhouse

import (
	"encoding/json"
	"time"
)

type Milestone struct {
	Categories          []Category `json:"categories"`
	Completed           bool       `json:"completed"`
	CompletedAt         time.Time  `json:"completed_at"`
	CompletedAtOverride time.Time  `json:"completed_at_override"`
	CreatedAt           time.Time  `json:"created_at"`
	Description         string     `json:"description"`
	EntityType          string     `json:"entity_type"`
	ID                  int64      `json:"id"`
	Name                string     `json:"name"`
	Position            int64      `json:"position"`
	Started             bool       `json:"started"`
	StartedAt           time.Time  `json:"started_at"`
	StartedAtOverride   time.Time  `json:"started_at_override"`
	State               string     `json:"state"`
	UpdatedAt           time.Time  `json:"updated_at"`
}

func (ch *Clubhouse) ListMilestones() ([]Milestone, error) {
	body, err := ch.listResources("milestones")
	if err != nil {
		return []Milestone{}, err
	}
	milestones := []Milestone{}
	json.Unmarshal(body, &milestones)
	return milestones, nil
}
