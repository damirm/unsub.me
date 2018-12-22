package subscription

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

var socialNetworks = make(map[string]SocialNetwork)

// Subscription represents any social "subscription".
// For example:
// - instagram profile
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
	SocialNetwork     string
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
func RegisterSocialNetwork(sns ...SocialNetwork) {
	for _, sn := range sns {
		socialNetworks[sn.Name()] = sn
	}
}

// List returns filtered list of subscriptions.
func List(filter Filter) (result []Subscription, err error) {
	var sns map[string]SocialNetwork

	if filter.SocialNetwork != "" {
		sn, ok := socialNetworks[filter.SocialNetwork]
		if !ok {
			return nil, fmt.Errorf("Unknown social network %s", filter.SocialNetwork)
		}
		sns[filter.SocialNetwork] = sn
	} else {
		sns = socialNetworks
	}

	for _, sn := range sns {
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
