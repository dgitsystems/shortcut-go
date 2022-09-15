package shortcut

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateGroup struct {
	Color         string      `json:"color,omitempty"`
	ColorKey      string      `json:"color_key,omitempty"`
	Description   string      `json:"description,omitempty"`
	DisplayIconID string      `json:"display_icon_id,omitempty"`
	MemberIds     []string    `json:"member_ids,omitempty"`
	MentionName   interface{} `json:"mention_name,omitempty"`
	Name          interface{} `json:"name,omitempty"`
	WorkflowIds   []int       `json:"workflow_ids,omitempty"`
}

type UpdateGroup struct {
	Archived      bool        `json:"archived,omitempty"`
	Color         string      `json:"color,omitempty"`
	ColorKey      string      `json:"color_key,omitempty"`
	Description   string      `json:"description,omitempty"`
	DisplayIconID string      `json:"display_icon_id,omitempty"`
	MemberIds     []string    `json:"member_ids,omitempty"`
	MentionName   interface{} `json:"mention_name,omitempty"`
	Name          interface{} `json:"name,omitempty"`
	WorkflowIds   []int       `json:"workflow_ids,omitempty"`
}

type Group struct {
	AppURL      string `json:"app_url,omitempty,omitempty"`
	Archived    bool   `json:"archived,omitempty,omitempty"`
	Color       string `json:"color,omitempty,omitempty"`
	ColorKey    string `json:"color_key,omitempty,omitempty"`
	Description string `json:"description,omitempty,omitempty"`
	DisplayIcon struct {
		CreatedAt  time.Time `json:"created_at,omitempty,omitempty"`
		EntityType string    `json:"entity_type,omitempty,omitempty"`
		ID         string    `json:"id,omitempty,omitempty"`
		UpdatedAt  time.Time `json:"updated_at,omitempty,omitempty"`
		URL        string    `json:"url,omitempty,omitempty"`
	} `json:"display_icon,omitempty,omitempty"`
	EntityType        string   `json:"entity_type,omitempty,omitempty"`
	ID                string   `json:"id,omitempty,omitempty"`
	MemberIds         []string `json:"member_ids,omitempty,omitempty"`
	MentionName       string   `json:"mention_name,omitempty,omitempty"`
	Name              string   `json:"name,omitempty,omitempty"`
	NumEpicsStarted   int      `json:"num_epics_started,omitempty,omitempty"`
	NumStories        int      `json:"num_stories,omitempty,omitempty"`
	NumStoriesStarted int      `json:"num_stories_started,omitempty,omitempty"`
	WorkflowIds       []int    `json:"workflow_ids,omitempty,omitempty"`
}

func (ch *Shortcut) ListGroups() ([]Group, error) {
	body, err := ch.listResources("groups")
	if err != nil {
		return []Group{}, err
	}
	groups := []Group{}
	json.Unmarshal(body, &groups)
	return groups, nil
}

func (ch *Shortcut) CreateGroup(newGroup CreateGroup) (Group, error) {
	jsonStr, _ := json.Marshal(newGroup)

	body, err := ch.createObject("groups", jsonStr)
	if err != nil {
		return Group{}, err
	}
	group := Group{}
	json.Unmarshal(body, &group)
	return group, nil
}

func (ch *Shortcut) GetGroup(projectID int64) (Group, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "groups", projectID))
	if err != nil {
		return Group{}, err
	}
	group := Group{}
	json.Unmarshal(body, &group)
	return group, nil
}

func (ch *Shortcut) UpdateGroup(updatedGroup UpdateGroup, projectID int64) (Group, error) {
	jsonStr, _ := json.Marshal(updatedGroup)
	body, err := ch.updateResource(fmt.Sprintf("%s/%d", "groups", projectID), jsonStr)
	if err != nil {
		return Group{}, err
	}
	group := Group{}
	json.Unmarshal(body, &group)
	return group, nil
}

func (ch *Shortcut) DeleteGroup(projectID int64) error {
	return ch.deleteResource(fmt.Sprintf("%s/%d", "groups", projectID))
}
