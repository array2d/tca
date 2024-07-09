package shell

import (
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func BashFile(script string) (exitCode int, stdouterr string, err error) {
	tmpFile, err := os.CreateTemp("", "script-*.sh")
	if err != nil {
		log.WithError(err).Errorln("create tmp failed")
	}
	defer os.Remove(tmpFile.Name()) // 程序退出时删除临时文件

	if _, err = tmpFile.Write([]byte(script)); err != nil {
		log.WithError(err).Errorln("write tmp failed")
	}
	if err = tmpFile.Close(); err != nil {
		log.WithError(err).Errorln("tmp failed")
	}
	// 为临时文件添加执行权限
	if err = os.Chmod(tmpFile.Name(), 0755); err != nil {
		log.WithError(err).Errorln("tmp failed")
	}

	// 执行脚本

	return ShellResult([]string{"/bin/bash", tmpFile.Name()}, nil)
}
func BashC(sh string) (exitCode int, stdouterr string, err error) {
	cmd := []string{"/bin/bash", "-c", sh}
	return ShellResult(cmd, nil)

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
