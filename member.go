package shortcut

import (
	"encoding/json"
	"fmt"
	"time"
)

type Member struct {
	ID                   string        `json:"id"`
	CreatedAt            time.Time     `json:"created_at"`
	CreatedWithoutInvite bool          `json:"created_without_invite"`
	Disabled             bool          `json:"disabled"`
	GroupIds             []interface{} `json:"group_ids"`
	Profile              Profile       `json:"profile"`
	Role                 string        `json:"role"`
	State                string        `json:"state"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

type Profile struct {
	Deactivated            bool   `json:"deactivated"`
	EmailAddress           string `json:"email_address"`
	GravatarHash           string `json:"gravatar_hash"`
	Membername             string `json:"membername"`
	MentionName            string `json:"mention_name"`
	Name                   string `json:"name"`
	TwoFactorAuthActivated bool   `json:"two_factor_auth_activated"`
}

func (ch *Shortcut) GetMember(memberID int64) (Member, error) {
	body, err := ch.getResource(fmt.Sprintf("%s/%d", "members", memberID))
	if err != nil {
		return Member{}, err
	}

	member := Member{}
	json.Unmarshal(body, &member)

	return member, nil
}

func (ch *Shortcut) ListMembers() ([]Member, error) {
	body, err := ch.listResources("members")
	if err != nil {
		return []Member{}, err
	}
	members := []Member{}
	json.Unmarshal(body, &members)
	return members, nil
}
