package dirs

import (
	"path/filepath"

	log "github.com/TeoDev1611/chevify/internal/utils/logger"
	"github.com/kirsle/configdir"
)

/// ConfigFile     string
/// LogFile        string
/// BaseConfigPath string
/// AliasesPath    string

var ConfigPath = configdir.LocalConfig("chevify")

func Directories() map[string]string {
	err := configdir.MakePath(ConfigPath) // Ensure it exists.
	if err != nil {
		log.Fatal(err.Error())
	}
	return map[string]string{
		"LogFile":    log.LogPath(),
		"ConfigFile": filepath.Join(ConfigPath, "settings.json"),
	}
}
