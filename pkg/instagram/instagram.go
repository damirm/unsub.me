package instagram

import (
	"net/url"
	"time"

	"github.com/damirm/unsub.me/pkg/subscription"
)

// Instagram represents instagram social network.
type Instagram struct {
	OAuthToken string
	ClientID   string
}

// FetchSubscriptions implementation for instagram network.
func (i *Instagram) FetchSubscriptions() ([]subscription.Subscription, error) {
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
func (i *Instagram) Unsubscribe(s subscription.Subscription) error {
	return nil
}
