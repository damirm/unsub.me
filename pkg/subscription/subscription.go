package subscription

import (
	"net/url"
	"time"
)

var socialNetworks []SocialNetwork

// Subscription represents any social "subscription".
// For example:
// - instagram page
// - youtube channel
// - github repository
// - twitch channel
type Subscription struct {
	ID           string
	Name         string
	Description  string
	Image        *url.URL
	LastActivity time.Time
}

// SocialNetwork represents any social network with some kind of subscriptions.
type SocialNetwork interface {
	FetchSubscriptions() ([]Subscription, error)
	Unsubscribe(Subscription) error
}

// RegisterSocialNetwork registers social network driver.
func RegisterSocialNetwork(sn SocialNetwork) {
	socialNetworks = append(socialNetworks, sn)
}

// UnsubscribeAll unsubscribes from all subscriptions.
func UnsubscribeAll() error {
	for _, sn := range socialNetworks {
		subs, err := sn.FetchSubscriptions()
		if err != nil {
			return err
		}

		for _, s := range subs {
			err := sn.Unsubscribe(s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
