package version

import (
	"fmt"
)

// export build var
var (
	APP        string
	APPVERSION string
	GOVERSION  string
	TIME       string
	GIT        string
)

// String to string
func String() string {
	return fmt.Sprintf("app: %v\nversion: %v\ngoVersion: %v\nbuildTime: %v\nbuildTag: %v", APP, APPVERSION, GOVERSION, TIME, GIT)
}
