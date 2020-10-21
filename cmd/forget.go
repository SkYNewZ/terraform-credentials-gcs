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
	"github.com/SkYNewZ/terraform-credentials-gcs/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// forgetCmd represents the forget command
var forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Delete any stored token matching given hostname",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// First, read from file if exist
		log.Debugln("Reading file")
		tokens, err := internal.ReadTokenFile(tokenFilePath)
		if err != nil {
			return err
		}

		log.Debugf("Found %d token(s). Searching matching credentials for source [%s]", len(tokens), args[0])
		var c internal.TokenSources = make(internal.TokenSources, 0)
		for _, t := range tokens {
			if t.Source != args[0] {
				c = append(c, t)
			}
		}

		// If nothing changed
		if len(c) == len(tokens) {
			log.Debugln("Nothing to do")
			return nil
		}

		// Rewrite final file
		log.Debugf("Rewrite final files with %d token(s)", len(c))
		if err = c.WriteToFile(tokenFilePath); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(forgetCmd)
}
