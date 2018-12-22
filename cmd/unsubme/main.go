package main

import (
	"os"
	"time"

	"github.com/damirm/unsub.me/pkg/instagram"
	"github.com/damirm/unsub.me/pkg/subscription"
	"github.com/spf13/cobra"
)

var until *time.Duration

func main() {
	sn := &instagram.Instagram{OAuthToken: "..."}
	subscription.RegisterSocialNetwork(sn)

	unsubscribeAll := &cobra.Command{
		Use: "unsubscribe-all",
		Run: func(cmd *cobra.Command, args []string) {
			subscription.UnsubscribeAll()
		},
	}

	list := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			filter := subscription.Filter{
				LastActivityUntil: time.Now().Add(*until),
			}
			subscription.List(filter)
		},
	}

	until = list.Flags().Duration("until", 0, "")

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(unsubscribeAll)
	rootCmd.AddCommand(list)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
