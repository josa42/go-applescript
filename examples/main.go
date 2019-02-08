package main

import (
	"fmt"

	"github.com/josa42/go-applescript"
)

func main() {
	s, _ := applescript.Run(`
		tell application "Safari"
			set theURL to URL of current tab of window 1
		end tell
	`)
	fmt.Printf("|%s|\n", s)
}
