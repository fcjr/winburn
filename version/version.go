package version

import (
	"fmt"
)

var Version = "source"

func String() string {
	return fmt.Sprintf("winburn version: %s\n", Version)
}
