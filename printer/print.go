package printer

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func AddToPrintQueue(path string) error {
	// check path is file
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return errors.New("path is not a file")
	}
	cmd := exec.Command("lpr", path)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func IsPrintQueueEmpty() (bool, error) {
	cmd := exec.Command("lpq")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	if strings.Contains(string(out), "no entries") {
		return true, nil
	}
	return false, nil
}
