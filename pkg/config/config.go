package config

// Config represents some configuration.
type Config struct {
	Instagram InstagramConfig `json:"instagram"`
	Twitch    TwitchConfig    `json:"twitch"`
}

// TwitchConfig represents twich auth requirements.
type TwitchConfig struct {
	ClientSecret string `json:"client_secret"`
	ClientID     string `json:"client_id"`
}

// InstagramConfig represents instagram api requirement credentials.
type InstagramConfig struct {
	ClientSecret string `json:"client_secret"`
	ClientID     string `json:"client_id"`
}
