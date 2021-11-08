package shortcut

import (
	"encoding/json"
	"fmt"
	"time"
)

//CreateEpic is the object passed to Shortcut API to create an epic.
//Required fields are:
// CreateEpic.Name
type CreateEpic struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Description string    `json:"description,omitempty"`
	ExternalID  string    `json:"external_id,omitempty"`
	FollowerIds []string  `json:"follower_ids,omitempty"`
	Name        string    `json:"name"`
	OwnerIds    []string  `json:"owner_ids,omitempty"`
	State       string    `json:"state,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

//An Epic is a collection of stories that together might make up a release, a milestone, or some other large initiative that your organization is working on.
type Epic struct {
	Archived            bool              `json:"archived"`
	Comments            []ThreadedComment `json:"comments"`
	Completed           bool              `json:"completed"`
	CompletedAt         time.Time         `json:"completed_at"`
	CompletedAtOverride time.Time         `json:"completed_at_override"`
	CreatedAt           time.Time         `json:"created_at"`
	Deadline            time.Time         `json:"deadline"`
	Description         string            `json:"description"`
	EpicStateID         int64             `json:"epic_state_id"`
	ExternalID          int64             `json:"external_id"`
	FollowerIds         []string          `json:"follower_ids"`
	ID                  int64             `json:"id"`
	Labels              []Label           `json:"labels"`
	MentionIds          []string          `json:"mention_ids"`
	MilestoneID         int64             `json:"milestone_id"`
	Name                string            `json:"name"`
	OwnerIds            []string          `json:"owner_ids"`
	Position            int64             `json:"position"`
	ProjectIds          []int64           `json:"project_ids"`
	RequestedByID       string            `json:"requested_by_id"`
	Started             bool              `json:"started"`
	StartedAt           time.Time         `json:"started_at"`
	StartedAtOverride   time.Time         `json:"started_at_override"`
	State               string            `json:"state"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

// UpdateEpic the body used for ch.EpicUpdate()
type UpdateEpic struct {
	AfterID     int64     `json:"after_id,omitempty"`
	Archived    bool      `json:"archived,omitempty"`
	BeforeID    int64     `json:"before_id,omitempty"`
	Deadline    time.Time `json:"deadline,omitempty"`
	Description string    `json:"description,omitempty"`
	FollowerIds []string  `json:"follower_ids,omitempty"`
	Name        string    `json:"name,omitempty"`
	OwnerIds    []string  `json:"owner_ids,omitempty"`
	State       string    `json:"state,omitempty"`
}

// EpicGet returns information about the selected Epic.
//calls GET https://api.shortcut.io/api/v3/epics/{epicID} to retrieve the specified epicID
func (ch *Shortcut) EpicGet(epicID int64) (Epic, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "epics", epicID))
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

// EpicUpdate can be used to update numerous fields in the Epic. The only required parameter is Epic ID, which can be found in the Shortcut UI.
//calls PUT https://api.shortcut.io/api/v3/epics/{epicID} and updates it with the data in the UpdateEpic object.
func (ch *Shortcut) EpicUpdate(updatedEpic UpdateEpic, epicID int64) (Epic, error) {
	jsonStr, _ := json.Marshal(updatedEpic)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "epics", epicID), jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}

// EpicDelete can be used to delete the Epic. The only required parameter is Epic ID.
//Calls DELETE https://api.shortcut.io/api/v3/epics/{epicID}
func (ch *Shortcut) EpicDelete(epicID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "epics", epicID))
}

// EpicList returns a list of all Epics and their attributes.
//Calls GET https://api.shortcut.io/api/v3/epics/
func (ch *Shortcut) EpicList() ([]Epic, error) {
	body, err := ch.listResources("epics")
	if err != nil {
		return []Epic{}, err
	}
	epics := []Epic{}
	json.Unmarshal(body, &epics)
	return epics, nil
}

//EpicCreate allows you to create a new Epic in Shortcut.
//Calls POST https://api.shortcut.io/api/v3/epics/
func (ch *Shortcut) EpicCreate(newEpic CreateEpic) (Epic, error) {
	jsonStr, _ := json.Marshal(newEpic)
	body, err := ch.createObject("epics", jsonStr)
	if err != nil {
		return Epic{}, err
	}
	epic := Epic{}
	json.Unmarshal(body, &epic)
	return epic, nil
}
