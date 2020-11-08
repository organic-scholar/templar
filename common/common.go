package common

import (
	"os"

	"github.com/spf13/afero"
)

var memFs = afero.NewMemMapFs()

func GetFs() afero.Fs {
	if os.Getenv("GO_ENV") == "testing" {
		return memFs
	}
	return afero.NewOsFs()
}
