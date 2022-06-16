package browser

import (
	"os/exec"
	"runtime"

	log "github.com/TeoDev1611/chevify/internal/utils/logger"
)

func OpenUrlBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
		log.Info("Opened the url in the browser")
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		log.Info("Opened the url in the browser")
	case "darwin":
		err = exec.Command("open", url).Start()
		log.Info("Opened the url in the browser")

	default:
		log.Warn("Platform not supported")
	}
	if err != nil {
		log.Error(err.Error())
	}
}
