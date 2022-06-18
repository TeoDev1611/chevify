package install

import (
	"errors"

	"github.com/TeoDev1611/chevify/internal/utils/browser"
	// git "github.com/go-git/go-git/v5"
	log "github.com/TeoDev1611/chevify/internal/utils/logger"
	// "github.com/TeoDev1611/chevify/internal/core/dirs"
)

// "github.com/go-git/go-git/v5"

// Types for the Installation parameters
type Params struct {
	Name   string
	Branch string
	Editor string
	Speed  int
	Force  bool
}

// Error Variables
var (
	ErrMissingEditor = errors.New("Editor Field is Missing!")
)

// Default Constants
const (
	DefaultName   = "exampleConfigName"
	DefaultBranch = "main"
	Speed         = 999
	Force         = false
)

// Validator Of the Params
func (o *Params) Validate() error {
	if o.Editor == "" {
		return ErrMissingEditor
	}

	if o.Name == "" {
		o.Name = DefaultName
	}

	if o.Branch == "" {
		o.Branch = DefaultBranch
	}

	return nil
}

func InstallNewConfig(url string, o *Params) {
	if err := o.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	isUrl := browser.IsUrl(url)
	if !isUrl {
		log.Fatal("This is not a url valid!")
	}
}
