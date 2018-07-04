package writeas

import (
	"testing"
)

func TestGetPosts(t *testing.T) {
	dwac := NewDevClient()

	au, err := dwac.LogIn("writeas-testuser", "RGX-SQg-a7m-4pV")

	if err != nil {
		t.Errorf("failure to login: %s", err)
	}

	dwac.SetToken(au.AccessToken)

	p, err := dwac.GetPosts()

	if err != nil {
		t.Errorf("failure to get posts: %s", err)
	}

	if len(p) != 0 {
		t.Errorf("expected a length of 0 but got: %d", len(p))
	}
}
