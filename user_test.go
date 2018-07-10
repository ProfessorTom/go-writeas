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

func CreateMultiplePosts(c *Client) (*[]PostParams, error) {
	pp := &[]PostParams{
		PostParams{
			Title:   "Familar Song",
			Content: "This is the song that never ends.",
			Font:    "sans",
		},
		PostParams{
			Title:   "Going around rocks",
			Content: "Round the ragged rock, the ragged rascal ran.",
			Font:    "sans",
		},
		PostParams{
			Title:   "Girl on a Beach",
			Content: "She sells seashells by the seashore.",
			Font:    "sans",
		},
	}

	count := len(*pp)
	for i := 0; i < count; i++ {
		p, err := c.CreatePost(&(*pp)[i])

		if err != nil {
			return nil, err
		}

		// kludge to ignore returned post
		if p != nil {
			continue
		}
	}

	return pp, nil
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
