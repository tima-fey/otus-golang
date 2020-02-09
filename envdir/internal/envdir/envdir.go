package envdir

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func ReadDir(dir string) (map[string]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	answer := make(map[string]string)
	for _, fileInfo := range files {
		file, err := os.Open(dir + "/" + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		if strings.Contains(file.Name(), "=") {
			continue
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		if len(txtlines) > 0 {
			if strings.Contains(txtlines[0], "=") {
				continue
			}
			answer[fileInfo.Name()] = strings.TrimSpace(txtlines[0])
		} else {
			answer[fileInfo.Name()] = ""
		}
	}
	return answer, nil
}
func RunCmd(cmd []string, env map[string]string) int {
	for key, value := range env {
		if value == "" {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, value)
		}
	}
	commandExec := exec.Command(cmd[0], cmd[1:]...)
	commandExec.Stdout = os.Stdout
	commandExec.Stderr = os.Stderr
	err := commandExec.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if status, ok := exitError.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus()
			}
		}
	}
	return 0
}

func Envdir(dir string, cmd []string) (int, error) {
	envs, err := ReadDir(dir)
	if err != nil {
		return 0, err
	}
	exitCode := RunCmd(cmd, envs)
	return exitCode, nil
}
