package auth

import (
	"fmt"

	"github.com/damirm/unsub.me/pkg/subscription"
)

// Wizard provides step by step authorization.
type Wizard struct {
	SN subscription.SocialNetwork
}

// TerminalAuth provides terminal wizard authentication.
func (w *Wizard) TerminalAuth() error {
	authURL := w.SN.AuthURL()

	fmt.Printf(
		"Go to the following link in your browser then type the authorization code:\n%v\n",
		authURL.String(),
	)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		return nil
	}

	if err := w.SN.OnAuthCode(code); err != nil {
		return err
	}

	return nil
}
