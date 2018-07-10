package writeas

import (
	"bytes"
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

func GetPostParamsFromPostArray(p *[]Post) *[]PostParams {
	pp := make([]PostParams, len(*p))

	for i := 0; i < len(*p); i++ {
		params := PostParams{
			Title:   (*p)[i].Title,
			Content: (*p)[i].Content,
			Font:    (*p)[i].Font,
		}

		fmt.Printf("\nparams: %#v\n", params)
		pp[i] = params
	}
	return &pp
}

func ReverseSlice(epp *[]PostParams) {
	for i, j := 0, len(*epp)-1; i < j; i, j = i+1, j-1 {
		(*epp)[i], (*epp)[j] = (*epp)[j], (*epp)[i]
	}

	// fmt.Printf("reversed: %s", PrintSlice(epp))
}

func PrintSlice(pp *[]PostParams) string {
	var buffer bytes.Buffer

	for i := 0; i < len(*pp); i++ {
		buffer.WriteString("Title: ")
		buffer.WriteString((*pp)[i].Title + "\n")
		// buffer.WriteString("\n")
		buffer.WriteString("Content: ")
		buffer.WriteString((*pp)[i].Content + "\n")
		buffer.WriteString("\n")
	}

	return buffer.String()
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

func TestGetPostsWithPosts(t *testing.T) {
	dwac, err := DeleteAllPosts()
	if err != nil {
		t.Errorf("%s", err)
	}

	pp, err := CreateMultiplePosts(dwac)
	if err != nil {
		t.Errorf("%s", err)
	}

	p, err := dwac.GetPosts()
	if err != nil {
		t.Errorf("failure to get posts: %s", err)
	}

	if len(*p) != len(*pp) {
		t.Errorf("expected a length of %d but got a length of  %d", len(*pp), len(*p))
	}
}

func TestGetPostPostOrder(t *testing.T) {
	dwac, err := DeleteAllPosts()
	if err != nil {
		t.Errorf("%s", err)
	}

	fmt.Println(dwac.token)

	epp, err := CreateMultiplePosts(dwac)
	if err != nil {
		t.Errorf("%s", err)
	}

	posts, err := dwac.GetPosts()
	if err != nil {
		t.Errorf("failure to get posts: %s", err)
	}

	// fmt.Printf("epp: %s", PrintSlice(epp))
	ReverseSlice(epp)

	// fmt.Printf("epp: %s", PrintSlice(epp))
	app := GetPostParamsFromPostArray(posts)
	// fmt.Printf("app: %s", PrintSlice(app))

	if PrintSlice(app) != PrintSlice(epp) {
		t.Errorf("expected \n%s\nbut got %s\n", PrintSlice(epp), PrintSlice(app))
	}

}
