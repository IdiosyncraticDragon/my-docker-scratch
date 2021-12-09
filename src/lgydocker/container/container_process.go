package container

import (
	"syscall"
	"os/exec"
	"os"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init", command} //这里的command是命令行的原始参数，现在在所有参数之前加入一个init参数，用来触发容器的初始化过程。
	cmd := exec.Command("/proc/self/exe", args...)//'...'是go的一种语法糖。它的第一个用法主要是用于函数有多个不定参数的情况，可以接受多个不确定数量的参数。第二个用法是slice可以被打散进行传递。这里是第二种用法。
	cmd.SysProcAttr = &syscall.SysProcAttr{//这里来定义进程的namespace
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
