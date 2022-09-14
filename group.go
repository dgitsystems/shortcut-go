package shortcut

import (
	"encoding/json"
	"fmt"
	"time"
)

type CreateGroup struct {
	Color         string      `json:"color"`
	ColorKey      string      `json:"color_key"`
	Description   string      `json:"description"`
	DisplayIconID string      `json:"display_icon_id"`
	MemberIds     []string    `json:"member_ids"`
	MentionName   interface{} `json:"mention_name"`
	Name          interface{} `json:"name"`
	WorkflowIds   []int       `json:"workflow_ids"`
}

type UpdateGroup struct {
	Archived      bool        `json:"archived"`
	Color         string      `json:"color"`
	ColorKey      string      `json:"color_key"`
	Description   string      `json:"description"`
	DisplayIconID string      `json:"display_icon_id"`
	MemberIds     []string    `json:"member_ids"`
	MentionName   interface{} `json:"mention_name"`
	Name          interface{} `json:"name"`
	WorkflowIds   []int       `json:"workflow_ids"`
}

type Group struct {
	AppURL      string `json:"app_url,omitempty"`
	Archived    bool   `json:"archived,omitempty"`
	Color       string `json:"color,omitempty"`
	ColorKey    string `json:"color_key,omitempty"`
	Description string `json:"description,omitempty"`
	DisplayIcon struct {
		CreatedAt  time.Time `json:"created_at,omitempty"`
		EntityType string    `json:"entity_type,omitempty"`
		ID         string    `json:"id,omitempty"`
		UpdatedAt  time.Time `json:"updated_at,omitempty"`
		URL        string    `json:"url,omitempty"`
	} `json:"display_icon,omitempty"`
	EntityType        string   `json:"entity_type,omitempty"`
	ID                string   `json:"id,omitempty"`
	MemberIds         []string `json:"member_ids,omitempty"`
	MentionName       string   `json:"mention_name,omitempty"`
	Name              string   `json:"name,omitempty"`
	NumEpicsStarted   int      `json:"num_epics_started,omitempty"`
	NumStories        int      `json:"num_stories,omitempty"`
	NumStoriesStarted int      `json:"num_stories_started,omitempty"`
	WorkflowIds       []int    `json:"workflow_ids,omitempty"`
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
