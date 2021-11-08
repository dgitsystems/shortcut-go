package shortcut_test

import (
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strings"
	"testing"

	// Frameworks
	sh "github.com/dgitsystems/shortcut-go/v3"
)

// In order to run tests, you will need to put your
// API token in a file called .shortcut in your
// home directory

func Test_Milestones_000(t *testing.T) {
	if client := sh.New(""); client == nil {
		t.Error("Invalid client == nil")
	} else {
		t.Log(client)
	}
}

func Test_Milestones_001(t *testing.T) {
	if token, err := getMilestonesToken(); err != nil {
		t.Error(err)
	} else if client := sh.New(token).SetDebug(true); client == nil {
		t.Error("Invalid client == nil")
	} else if milestones, err := client.ListMilestones(); err != nil {
		t.Error(err)
	} else {
		t.Log(milestones)
	}
}

/////////////////////////////////////////////////////

func getMilestonesToken() (string, error) {
	if user, err := user.Current(); err != nil {
		return "", err
	} else if bytes, err := ioutil.ReadFile(filepath.Join(user.HomeDir, ".shortcut")); err != nil {
		return "", err
	} else {
		return strings.TrimSpace(string(bytes)), nil
	}
}
