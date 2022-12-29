package misc

import (
	"fmt"
	"os"
	"path"
)

var (
	Software       = "Kubeshark"
	Program        = "kubeshark"
	Website        = "https://kubeshark.co"
	Ver            = "0.0"
	Branch         = "develop"
	GitCommitHash  = "" // this var is overridden using ldflags in makefile when building
	BuildTimestamp = "" // this var is overridden using ldflags in makefile when building
	RBACVersion    = "v1"
	Platform       = ""
)

func GetDotFolderPath() string {
	home, homeDirErr := os.UserHomeDir()
	if homeDirErr != nil {
		return ""
	}
	return path.Join(home, fmt.Sprintf(".%s", Program))
}