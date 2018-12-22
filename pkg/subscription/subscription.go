package subscription

import (
	"log"
	"net/url"
	"strings"
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

// Filter represents various of filters.
type Filter struct {
	LastActivityUntil time.Time
	Name              string
	Categories        []string
}

func (f *Filter) passes(s Subscription) bool {
	if s.LastActivity.After(f.LastActivityUntil) {
		return false
	}

	if strings.TrimSpace(f.Name) != "" {
		sName := strings.ToLower(strings.TrimSpace(s.Name))
		fName := strings.ToLower(strings.TrimSpace(f.Name))
		if !strings.Contains(sName, fName) {
			return false
		}
	}

	return true
}

// SocialNetwork represents any social network with some kind of subscriptions.
type SocialNetwork interface {
	Name() string
	FetchSubscriptions() ([]Subscription, error)
	Unsubscribe(Subscription) error
}

// RegisterSocialNetwork registers social network driver.
func RegisterSocialNetwork(sn SocialNetwork) {
	socialNetworks = append(socialNetworks, sn)
}

// List returns filtered list of subscriptions.
func List(filter Filter) (result []Subscription, err error) {
	for _, sn := range socialNetworks {
		subs, err := sn.FetchSubscriptions()
		if err != nil {
			return nil, err
		}
		for _, s := range subs {
			passes := filter.passes(s)
			if passes {
				result = append(result, s)
			}
		}
	}
	return
}

// UnsubscribeAll unsubscribes from all subscriptions.
func UnsubscribeAll() error {
	for _, sn := range socialNetworks {
		subs, err := sn.FetchSubscriptions()
		if err != nil {
			return err
		}

		log.Printf("Got %d subscriptions from %s social network\n", len(subs), sn.Name())

		for _, s := range subs {
			err := sn.Unsubscribe(s)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
