package shell

import (
	"os/exec"
)

func BashC(sh string) (exitCode int, stdouterr string, err error) {
	cmd := []string{"/bin/bash", "-c", sh}
	return ExecuteCommand(cmd, nil)
	return
}
func ExecuteCommand(cmds, envs []string) (exitCode int, stdouterr string, err error) {
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
