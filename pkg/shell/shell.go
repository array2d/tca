package shell

import (
	"os/exec"
)

func ExecuteCommand(cmds string, envs []string) (exitCode int, stdouterr string, err error) {
	cmd := exec.Command(cmds)
	// 设置环境变量
	cmd.Env = envs
	// 获取命令输出
	var std []byte
	std, err = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()
	stdouterr = string(std)
	return
}
