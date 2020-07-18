package version

import (
	"fmt"
	"runtime"
)

func Printversion() string {
	return fmt.Sprintf("%s", runtime.Version())
}
