package config

import (
	log "github.com/TeoDev1611/chevify/internal/utils/logger"
	"github.com/kirsle/configdir"
)

var ConfigPath = configdir.LocalConfig("chevify")

func MakeConfigFolder() {
	err := configdir.MakePath(ConfigPath) // Ensure it exists.
	if err != nil {
		log.Fatal(err.Error())
	}
}
