package shell

import (
	"os"
	"os/exec"
)

func BashC(sh string) (exitCode int, stdouterr string, err error) {
	cmd := []string{"/bin/bash", "-c", sh}
	return ShellResult(cmd, nil)
	return
}
func ShellStd(cmds, envs []string) (exitCode int, stdouterr string, err error) {
	cmd := exec.Command(cmds[0], cmds[1:]...)
	// 设置环境变量
	cmd.Env = envs
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 获取命令输出
	err = cmd.Run()
	exitCode = cmd.ProcessState.ExitCode()
	return
}

func ShellResult(cmds, envs []string) (exitCode int, stdouterr string, err error) {
	cmd := exec.Command(cmds[0], cmds[1:]...)
	// 设置环境变量
	cmd.Env = envs
	// 获取命令输出
	var std []byte
	std, err = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()
	stdouterr = string(std)
	return
}
