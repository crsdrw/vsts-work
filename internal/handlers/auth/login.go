package auth

import (
	"log"
	"github.com/crsdrw/vsts-work/internal/flags"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

const user = "PAT token"

// Login invokes the login handler
func Login(cmd *cobra.Command, args []string) error {

    // set password
    err := keyring.Set(flags.Instance, user, flags.Token)
    if err != nil {
        log.Fatal(err)
    }

	return nil
}
