/*
User Namesapce主要用来隔离用户的用户组ID。一个进程的User ID和Group ID在User Namesapce内外是不同的。
*/
package main

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main(){
	cmd := exec.Command("sh") // 指定被fork出来的新进程内的初始命令
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER, //使用CLONE_NEWUTS标识来创建一个UTC namesapce，使用CLONE_NEWIPC表示来创建IPC namesapce,使用CLONE_NEWPID标识来创建PID namespace
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {//go封装了对于系统clone()函数的调用，这段代码执行后会进入一个sh运行环境中
		log.Fatal(err)
	}
}
