package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
)

// ReadTokenFile read the current token from file if exist
func ReadTokenFile(tokenFilePath string) (TokenSources, error) {
	// File does not exist
	if _, err := os.Stat(tokenFilePath); os.IsNotExist(err) {
		log.Debugln("File does not exist, skip reading")
		return nil, nil
	}

	data, err := ioutil.ReadFile(tokenFilePath)
	if err != nil {
		return nil, err
	}

	// File exist but empty
	if len(data) == 0 {
		return nil, nil
	}

	var tokens []TokenSource
	err = json.Unmarshal(data, &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
