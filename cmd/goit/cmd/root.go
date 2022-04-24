package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pkg/browser"
	"github.com/pojntfx/goit/pkg/token"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	errMissingOIDCIssuer      = errors.New("missing OIDC issuer")
	errMissingOIDCClientID    = errors.New("missing OIDC client ID")
	errMissingOIDCRedirectURL = errors.New("missing OIDC redirect URL")
)

const (
	oidcRedirectURLFlag = "oidc-redirect-url"
	oidcIssuerFlag      = "oidc-issuer"
	oidcClientIDFlag    = "oidc-client-id"
	verboseFlag         = "verbose"
)

var rootCmd = &cobra.Command{
	Use:   "goit",
	Short: "Get a OIDC token from your terminal",
	Long: `Get a OIDC token from your terminal.

Find more information at:
https://github.com/pojntfx/goit`,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("")
		viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

		if err := viper.BindPFlags(cmd.PersistentFlags()); err != nil {
			return err
		}

		if strings.TrimSpace(viper.GetString(oidcIssuerFlag)) == "" {
			return errMissingOIDCIssuer
		}

		if strings.TrimSpace(viper.GetString(oidcClientIDFlag)) == "" {
			return errMissingOIDCClientID
		}

		if strings.TrimSpace(viper.GetString(oidcRedirectURLFlag)) == "" {
			return errMissingOIDCRedirectURL
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		manager := token.NewTokenManager(
			viper.GetString(oidcIssuerFlag),
			viper.GetString(oidcClientIDFlag),
			viper.GetString(oidcRedirectURLFlag),

			func(s string) error {
				// Don't write any output to stdout except for the token
				browser.Stderr = os.Stderr
				browser.Stdout = os.Stderr

				if err := browser.OpenURL(s); err != nil {
					log.Printf(`Could not open browser, please open the following URL in your browser manually to authorize:
%v`, s)
				}

				return nil
			},

			ctx,
		)

		token, err := manager.GetToken()
		if err != nil {
			return err
		}

		log.Println("Successfully got the following OIDC access token:")

		fmt.Println(token)

		return nil
	},
}

func Execute() error {
	rootCmd.PersistentFlags().BoolP(verboseFlag, "v", false, "Enable verbose logging")
	rootCmd.PersistentFlags().String(oidcIssuerFlag, "", "OIDC Issuer (i.e. https://pojntfx.eu.auth0.com/) (can also be set using the OIDC_ISSUER env variable)")
	rootCmd.PersistentFlags().String(oidcClientIDFlag, "", "OIDC Client ID (i.e. myoidcclientid) (can also be set using the OIDC_CLIENT_ID env variable)")
	rootCmd.PersistentFlags().String(oidcRedirectURLFlag, "http://localhost:11337", "OIDC Redirect URL")

	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		return err
	}

	viper.AutomaticEnv()

	return rootCmd.Execute()
}
