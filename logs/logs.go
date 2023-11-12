package logs

import (
	"fmt"
	"os"
)

func Info(format string, a ...any) {
	fmt.Printf("[INFO] %s\n", fmt.Sprintf(format, a...))
}

func Error(format string, a ...any) {
	fmt.Printf("[ERROR] %s\n", fmt.Sprintf(format, a...))
	os.Exit(1)
}
