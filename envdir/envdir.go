package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func ReadDir(dir string) (map[string]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	answer := make(map[string]string)
	for _, fileInfo := range files {
		fmt.Println(fileInfo.Name())
		file, err := os.Open(dir + "/" + fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var txtlines []string
		for scanner.Scan() {
			txtlines = append(txtlines, scanner.Text())
		}
		if len(txtlines) > 0 {
			answer[fileInfo.Name()] = txtlines[0]
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
			return exitError.ExitCode()
		}
	}
	return 0
}

func envdir(dir string, cmd []string) (int, error) {
	envs, err := ReadDir(dir)
	if err != nil {
		return 0, err
	}
	exitCode := RunCmd(cmd, envs)
	return exitCode, nil
}
func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error. Specify env dir and command to execute")
		os.Exit(1)
	}
	exitCode, err := envdir(os.Args[1], os.Args[2:])
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitCode)
}
