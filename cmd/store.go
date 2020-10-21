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
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/SkYNewZ/terraform-credentials-gcs/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store given token at given hostname",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Test if something on stdin
		stat, err := os.Stdin.Stat()
		if err != nil {
			return err
		}

		if (stat.Mode() & os.ModeCharDevice) != 0 {
			return errors.New("The command is intended to work with pipes")
		}

		// Read stdin and parse it
		str, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		var t internal.TerraformToken
		err = json.Unmarshal(str, &t)
		if err != nil {
			return err
		}
		log.Debugf("Receive token source [%s]", args[0])

		// Append it to file
		// First, read from file if exist
		log.Debugln("Reading file")
		tokens, err := internal.ReadTokenFile(tokenFilePath)
		if err != nil {
			return err
		}

		for _, token := range tokens {
			// Test not exist already
			if token.Source == args[0] {
				return errors.New("Token for hostname already configured")
			}
		}

		log.Debugf("Found %d token(s). Inserting [%s]", len(tokens), args[0])
		tokens = append(tokens, internal.TokenSource{
			Source: args[0],
			Token:  t.Token,
		})

		log.Debugf("Writing %d token(s) to file", len(tokens))
		if err = tokens.WriteToFile(tokenFilePath); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}
