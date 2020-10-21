/*
Copyright © 2020 Quentin Lemaire <quentin@lemairepro.fr>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/SkYNewZ/terraform-credentials-gcs/internal"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
// Search if token already saved in a file and returns its. Else, ask Google a new token
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a Google credentials token",
	Long:  "Return a token if present in file matching given hostname, else ask Google Default Credentials",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// First, read from file if exist
		log.Debugln("Searching token(s) from file")
		tokens, err := internal.ReadTokenFile(tokenFilePath)
		if err != nil {
			return err
		}

		// If file exist and non empty
		if tokens != nil {
			log.Debugf("Found %d token(s). Searching matching credentials for source [%s]", len(tokens), args[0])
			for _, t := range tokens {
				if t.Source == args[0] {
					log.Debugf("Token for [%s] found from file", t.Source)
					// Convert to Terraform JSON object
					o := internal.TerraformToken{Token: t.Token}
					fmt.Println(o.JSON())
					return nil
				}
			}
		}

		// Else, ask Google a new token
		log.Debugln("Ask Google for access token")
		t, err := internal.GetGoogleAccessToken(ctx)
		if err != nil {
			return err
		}

		log.Debugln("Received token")
		o := internal.TerraformToken{Token: t}
		fmt.Println(o.JSON())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
