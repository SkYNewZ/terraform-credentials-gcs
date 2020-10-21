package internal

import (
	"encoding/json"
	"io/ioutil"
)

// TokenSource object to save in file
type TokenSource struct {
	Source string `json:"source"`
	Token  string `json:"token"`
}

// TokenSources multiple TokenSource
type TokenSources []TokenSource

// TerraformToken JSON token object
type TerraformToken struct {
	Token string `json:"token"`
}

// JSON return current token instance as JSON bytes
func (t *TerraformToken) JSON() string {
	s, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(s)
}

// JSON return current token instance as JSON bytes
func (t *TokenSources) JSON() string {
	s, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(s)
}

// WriteToFile write tokens to given file
func (t *TokenSources) WriteToFile(filepath string) error {
	d, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, d, 0644)
}
