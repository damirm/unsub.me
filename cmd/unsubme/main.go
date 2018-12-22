package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/damirm/unsub.me/cmd/unsubme/list"
	"github.com/damirm/unsub.me/pkg/config"
	"github.com/damirm/unsub.me/pkg/instagram"
	"github.com/damirm/unsub.me/pkg/subscription"
	"github.com/damirm/unsub.me/pkg/twitch"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPath *string
var cfg config.Config

func registerAllSocialNetworks() {
	inst := &instagram.Instagram{
		ClientID:     cfg.Instagram.ClientID,
		ClientSecret: cfg.Instagram.ClientSecret,
	}
	twi := &twitch.Twitch{
		ClientID:     cfg.Twitch.ClientID,
		ClientSecret: cfg.Twitch.ClientSecret,
	}

	subscription.RegisterSocialNetwork(inst, twi)
}

func main() {
	rootCmd := &cobra.Command{
		Use: "unsubme",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			viper.SetConfigFile(*configPath)
			viper.SetConfigType("yaml")
			err := viper.ReadInConfig()
			if err != nil {
				return err
			}
			err = viper.Unmarshal(&cfg, func(decoderConfig *mapstructure.DecoderConfig) {
				decoderConfig.TagName = "json"
			})
			if err != nil {
				return err
			}

			registerAllSocialNetworks()

			return nil
		},
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	defaultConfigPath := filepath.Join(usr.HomeDir, ".unsubme.yaml")
	configPath = rootCmd.PersistentFlags().StringP("config", "c", defaultConfigPath, "")

	rootCmd.AddCommand(list.Command(&cfg))

	err = rootCmd.Execute()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
