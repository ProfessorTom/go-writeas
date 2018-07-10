package writeas

import (
	"fmt"
	"testing"
)

const (
	devUN = "writeas-testuser"
	devPW = "RGX-SQg-a7m-4pV"
)

func CreateSessionForTestUser() (*Client, error) {
	dwac := NewDevClient()

	au, err := dwac.LogIn(devUN, devPW)

	if err != nil {
		return nil, fmt.Errorf("failure to login: %s", err)
	}

	dwac.SetToken(au.AccessToken)
	return dwac, nil
}

func DeleteAllPosts() (*Client, error) {
	dwac, err := CreateSessionForTestUser()
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	p, err := dwac.GetPosts()
	if err != nil {
		return nil, fmt.Errorf("failure to get posts: %s", err)
	}

	count := len(*p)
	fmt.Printf("\nnumber of posts: %d\n\n", len(*p))
	if count != 0 {
		for i := 0; i < count; i++ {
			// Delete post
			err = dwac.DeletePost(&PostParams{
				OwnedPostParams: OwnedPostParams{
					ID:    (*p)[i].ID,
					Token: dwac.token,
				},
			})
			if err != nil {
				return nil, fmt.Errorf("Post delete failed: %v", err)
			}
		}
	}

	return dwac, nil
}

func TestGetPostsWithNoPosts(t *testing.T) {

	dwac, err := DeleteAllPosts()
	if err != nil {
		t.Errorf("%s", err)
	}

	p, err := dwac.GetPosts()
	if err != nil {
		t.Errorf("failure to get posts: %s", err)
	}

	if len(*p) != 0 {
		t.Errorf("expected a length of 0 but got: %d", len(*p))
	}
}
