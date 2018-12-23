package twitch

import (
	"net/url"
	"time"

	"github.com/damirm/unsub.me/pkg/subscription"
)

// Twitch represents instagram social network.
type Twitch struct {
	ClientID     string
	ClientSecret string
}

func (i *Twitch) Init() error {
	return nil
}

// Name return social network name.
func (i *Twitch) Name() string {
	return "Twitch"
}

// FetchSubscriptions implementation for instagram network.
func (i *Twitch) FetchSubscriptions() ([]subscription.Subscription, error) {
	image, err := url.Parse("https://raw.githubusercontent.com/ahmdrz/goinsta/v1/resources/goinsta-image.png")
	if err != nil {
		return nil, err
	}

	return []subscription.Subscription{
		{
			ID:           "1",
			Name:         "name",
			Description:  "description",
			Image:        image,
			LastActivity: time.Now(),
		},
		{
			ID:           "2",
			Name:         "name 2",
			Description:  "description 2",
			Image:        image,
			LastActivity: time.Now().Add(-1 * time.Hour),
		},
		{
			ID:           "3",
			Name:         "name 3",
			Description:  "description 3",
			Image:        image,
			LastActivity: time.Now().Add(-1 * 2 * time.Hour),
		},
	}, nil
}

// Unsubscribe from instagram channel
func (i *Twitch) Unsubscribe(s subscription.Subscription) error {
	return nil
}

// AuthURL returns twitch oauth url.
func (i *Twitch) AuthURL() *url.URL {
	res, _ := url.Parse("")
	return res
}

func (y *Twitch) OnAuthCode(code string) error {
	return nil
}
