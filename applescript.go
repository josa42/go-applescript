package applescript

import (
	"io"
	"os/exec"
	"strings"
)

// Run applacript
func Run(script string) (string, error) {
	return command("osascript")(script)
}

// RunScript :
func RunScript(path string) (string, error) {
	return command("osascript", path)()
}

func command(name string, args ...string) func(...string) (string, error) {
	return func(input ...string) (string, error) {
		cmd := exec.Command(name, args...)

		// cmd.Stderr = os.Stderr
		// cmd.Stdout = os.Stdout

		stdin, err := cmd.StdinPipe()
		if err != nil {
			return "", err
		}

		io.WriteString(stdin, strings.Join(input, " "))
		stdin.Close()

		b, err := cmd.Output()

		return strings.TrimSpace(string(b)), err
	}
}
