package list

import (
	"time"

	"github.com/damirm/unsub.me/pkg/config"
	"github.com/damirm/unsub.me/pkg/subscription"
	"github.com/spf13/cobra"
)

var (
	until         *time.Duration
	socialNetwork *string
)

// Command is a "list" command.
func Command(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: func(cmd *cobra.Command, args []string) {
			filter := subscription.Filter{
				LastActivityUntil: time.Now().Add(*until),
				SocialNetwork:     *socialNetwork,
			}
			subscription.List(filter)
		},
	}

	until = cmd.Flags().DurationP("until", "u", 0, "")
	socialNetwork = cmd.Flags().StringP("social-network", "s", "instagram", "")

	return cmd
}
