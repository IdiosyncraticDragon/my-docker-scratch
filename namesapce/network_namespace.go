/*
Network Namespace是用来个隔离网络设备、IP地址端口等网络栈的Namespace。Network Namespace可以让每个容器拥有自己独立的（虚拟的）网络设备，而且容器内的应用可以绑定到自己的端口，每个Namespace内的端口都不会互相冲突。在宿主机上搭建网桥后，就能实现容器间的通信。在不同的容器上，应用可以使用相同的端口。

CLONE_NEWNET标识符。
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET, //使用CLONE_NEWUTS标识来创建一个UTC namesapce，使用CLONE_NEWIPC表示来创建IPC namesapce,使用CLONE_NEWPID标识来创建PID namespace
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {//go封装了对于系统clone()函数的调用，这段代码执行后会进入一个sh运行环境中
		log.Fatal(err)
	}
}