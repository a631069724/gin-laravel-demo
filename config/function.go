package Config

import (
	"os"
	"os/exec"
	"path/filepath"
)

func ExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	return filepath.Abs(file)
}
