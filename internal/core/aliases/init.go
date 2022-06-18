package aliases

import (
	"path"
	"path/filepath"

	directories "github.com/TeoDev1611/chevify/internal/core/dirs"
	log "github.com/TeoDev1611/chevify/internal/utils/logger"
	missOs "github.com/TeoDev1611/chevify/internal/utils/os"
)

func GenerateAlias(editor, url string) {
	os, err := missOs.System()
	if err != nil {
		log.Fatal(err.Error())
	}
	dirs := directories.Directories()
	aliasName := path.Base(url)
	// Generate the Text From the editor
	if os == "windows" {
		e, err := missOs.Exists(filepath.Join(dirs["AliasesPath"], aliasName))
		if err != nil {
			log.Fatal(err.Error())
		}
		if e {
			log.Info("exists !")
		} else {
			log.Fatal("Not Found the Directory of the config for generate!")
		}
	}
}
