package commands

import (
	"github.com/crsdrw/vsts-work/internal/flags"
	"github.com/crsdrw/vsts-work/internal/handlers/auth"
	"github.com/spf13/cobra"
)

const loginShortString = "Login to VSTS"
const loginLongString = `
Login to vsts
`

var login = &cobra.Command{
	Use:   "login",
	Short: loginShortString,
	Long:  loginLongString,
	RunE:  auth.Login,
}

func init() {
	login.Flags().StringVarP(&flags.Instance, "instance", "i", "", "Instance of VSTS to connect to")
	login.Flags().StringVarP(&flags.Instance, "token", "t", "", "PAT token to use to connnect to VSTS")
	rootCmd.AddCommand(login)
}
