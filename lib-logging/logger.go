package logger

import (
	"fmt"

	"github.com/Boot-Error/golang-bazel-learn/lib-logging/constants"
)

func LogInfo(message string) {
	fmt.Printf("[%s] %s", constants.Info, message)
}
