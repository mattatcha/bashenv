package bashenv

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Source(filepath string) error {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	cmdStr := fmt.Sprintf("source %s 1>&2; env", filepath)

	out, err := exec.Command("bash", "-c", cmdStr).Output()
	if err != nil {
		return err
	}

	fileStr := string(file)
	outLines := strings.Split(string(out), "\n")

	for _, line := range outLines {
		lineSplit := strings.SplitN(line, "=", 2)
		if len(lineSplit) != 2 {
			continue
		}
		if strings.Contains(fileStr, lineSplit[0]) {
			os.Setenv(lineSplit[0], lineSplit[1])
		}
	}

	return nil
}
