package youtube

import (
	"context"
	"io/ioutil"
	"net/url"

	"github.com/damirm/unsub.me/pkg/subscription"

	"golang.org/x/oauth2/google"

	"golang.org/x/oauth2"
	gyoutube "google.golang.org/api/youtube/v3"
)

type Youtube struct {
	CredentialsFile string

	config  *oauth2.Config
	service *gyoutube.Service
}

func (y *Youtube) Init() error {
	b, err := ioutil.ReadFile(y.CredentialsFile)
	if err != nil {
		return err
	}
	y.config, err = google.ConfigFromJSON(b, gyoutube.YoutubeReadonlyScope)
	if err != nil {
		return err
	}
	return nil
}

func (y *Youtube) Name() string {
	return "Youtube"
}

func (y *Youtube) FetchSubscriptions() (result []subscription.Subscription, err error) {
	call := y.service.Subscriptions.List("snippet,contentDetails")
	response, err := call.Do()
	if err != nil {
		return nil, err
	}

	for _, item := range response.Items {
		snippet := item.Snippet
		image, err := url.Parse(snippet.Thumbnails.Default.Url)
		if err != nil {
			return nil, err
		}
		s := subscription.Subscription{
			ID:          snippet.ChannelId,
			Name:        snippet.ChannelTitle,
			Description: snippet.Description,
			Image:       image,
		}
		result = append(result, s)
	}
	return
}

// Unsubscribe from youtube channel
func (i *Youtube) Unsubscribe(s subscription.Subscription) error {
	return nil
}

func (y *Youtube) AuthURL() *url.URL {
	res, _ := url.Parse(y.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline))
	return res
}

func (y *Youtube) OnAuthCode(code string) error {
	token, err := y.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return err
	}
	client := y.config.Client(context.Background(), token)
	y.service, err = gyoutube.New(client)
	if err != nil {
		return err
	}
	return nil
}
