package utils

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

func GetPdfPageNum(path string) (int, error) {
	cmd := exec.Command("pdfinfo", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	re := regexp.MustCompile(`Pages:\s+(\d+)`)
	match := re.FindStringSubmatch(string(out))
	if len(match) < 2 {
		return 0, errors.New("can not get page num")
	}
	pageNum, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, err
	}
	return pageNum, nil
}
