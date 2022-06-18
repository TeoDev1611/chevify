package dirs

import (
	"path/filepath"

	"github.com/TeoDev1611/chevify/internal/core/config"
	log "github.com/TeoDev1611/chevify/internal/utils/logger"
)

/// ConfigFile     string
/// LogFile        string
/// BaseConfigPath string
/// AliasesPath    string

func Directories() map[string]string {
	config.MakeConfigFolder()
	return map[string]string{
		"LogFile":        log.LogPath(),
		"ConfigFile":     filepath.Join(config.ConfigPath, "settings.json"),
		"BaseConfigPath": config.ConfigPath,
		"AliasesPath":    filepath.Join(config.ConfigPath, "aliases"),
	}
}
