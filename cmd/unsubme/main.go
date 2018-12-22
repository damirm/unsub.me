package main

import (
	"github.com/damirm/unsub.me/pkg/instagram"
	"github.com/damirm/unsub.me/pkg/subscription"
)

func main() {
	sn := &instagram.Instagram{OAuthToken: "..."}

	subscription.RegisterSocialNetwork(sn)
	subscription.UnsubscribeAll()
}
