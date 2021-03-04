/*
mount namespace用来隔离各个进程看到的挂载点的视图。在不同的Mount Namespace的进程中，看到的文件系统的层次是不一样的。在Mount Namespace中调用mount()和unmount()仅仅只会影响当前Namespace内的文件系统，而对全局的文件系统没有影响。

这个功能非常类似linux听的chroot()，它也是将某一个子目录变为根节点。Mount Namespace也可以实现这个功能，且更加灵活和安全。

Mount Namespace是Linux第一个实现的Namespace类型，因为命名方式和之后的Namespace不一样（当时并没有规划那么多的Namespace）。标识是CLONE_NEWNS
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS, //使用CLONE_NEWUTS标识来创建一个UTC namesapce，使用CLONE_NEWIPC表示来创建IPC namesapce,使用CLONE_NEWPID标识来创建PID namespace
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {//go封装了对于系统clone()函数的调用，这段代码执行后会进入一个sh运行环境中
		log.Fatal(err)
	}
}