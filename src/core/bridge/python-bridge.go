package bridge

import (
	"fmt"
	"os/exec"
)

// CallPythonBridge runs a python script with given arguments and returns its output.
func CallPythonBridge(scriptPath string, args ...string) (string, error) {
	cmd := exec.Command("python", append([]string{scriptPath}, args...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("python error: %v, output: %s", err, string(out))
	}
	return string(out), nil
}
