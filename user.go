package writeas

import (
	"fmt"
	"time"
)

type (
	// AuthUser represents a just-authenticated user. It contains information
	// that'll only be returned once (now) per user session.
	AuthUser struct {
		AccessToken string `json:"access_token,omitempty"`
		Password    string `json:"password,omitempty"`
		User        *User  `json:"user"`
	}

	// User represents a registered Write.as user.
	User struct {
		Username string    `json:"username"`
		Email    string    `json:"email"`
		Created  time.Time `json:"created"`

		// Optional properties
		Subscription *UserSubscription `json:"subscription"`
	}

	// UserSubscription contains information about a user's Write.as
	// subscription.
	UserSubscription struct {
		Name       string    `json:"name"`
		Begin      time.Time `json:"begin"`
		End        time.Time `json:"end"`
		AutoRenew  bool      `json:"auto_renew"`
		Active     bool      `json:"is_active"`
		Delinquent bool      `json:"is_delinquent"`
	}
)

// GetPosts returns the posts for the currently authenticated user
// Authentication is stored in c.token which sets the Authorization header
func (c *Client) GetPosts() ([]Post, error) {
	var posts []Post

	env, err := c.get("me/posts", posts)

	if err != nil {
		return nil, err
	}

	var ok bool

	if posts, ok = env.Data.([]Post); !ok {
		return nil, fmt.Errorf("%v", posts)
	}

	return posts, nil

}
