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
	"context"
	"fmt"
	"os"
	"strings"

	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
)

var (
	tokenFilePath string
	ctx           context.Context

	rootCmd = &cobra.Command{
		Use:   "terraform-credentials-gcs",
		Short: "Terraform credentials helper for Google Cloud Storage",
		Long: `Terraform credentials helper used to pull privates providers from
	Google Cloud Storage. It will get, store and forget the default access token.
	
	You must provide either GOOGLE_APPLICATION_CREDENTIALS (service account file path)
	or GOOGLE_CREDENTIALS (service account content, trimmed).`,
	}

	// On-build tags
	version string = "dev"
	commit  string
	date    string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Version = fmt.Sprintf("%s - %s %s", version, date, commit)
}

func initConfig() {
	ctx = context.Background()

	// Init logger from TF_LOG
	switch strings.ToUpper(os.Getenv("TF_LOG")) {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&nested.Formatter{
		HideKeys:      true,
		ShowFullLevel: true,
	})

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln(err)
	}

	tokenFilePath = fmt.Sprintf("%s/.config/terraform-credentials-gcs", home)
	log.Debugf("File %s will be used", tokenFilePath)
}
