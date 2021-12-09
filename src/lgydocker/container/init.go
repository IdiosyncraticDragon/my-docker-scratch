package container

import (
	log "github.com/sirupsen/logrus"//这里把logrus这个库自定为log
	"syscall"
	"os"
)

/*
这是容器内执行的第一个进程.
docker创建一个容器后，会将容器的第一个进程指定为前台进程，它的进程id为1。
这个前台进程是不能kill的，因为它被kill的话，docker容器也就会退出。
而一般而言，这个id为1的进程都是容器初始化的init进程，而不是用户进程。
*/
func RunContainerInitProcess(command string, args []string) error {
	log.Infof("command rec while initiation -- %s", command)
	//格式示例：syscall.Mount("/dev/sda1", "/mnt1", "auto", 0, "w")
	syscall.Mount("", "/", "", uintptr(syscall.MS_PRIVATE | syscall.MS_REC), "")
	/* 
	MS_NOEXEC：在本文件系统不允许运行其他程序
	MS_NOSUID：在本文件系统中运行程序时，不允许set-user-ID或set-group-ID
	MS_NODEV：所有mount的系统都默认设定的参数
	*/
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	/*
	为什么要在容器中挂载/proc呢， 主要原因是因为ps、top等命令依赖于/proc目录。
	当隔离PID的时候，ps、top等命令还是未隔离的时候一样输出。 为了让隔离空间ps、top等命令只输出当前隔离空间的进程信息。需要单独挂载/proc目录。
	*/
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}
	/*
	syscall.Exec是调用Kernel的int execve(const char *filename, char *const argv[], char *const envp[])函数，这个函数会执行当前filename对应的程序（就是执行filename这个文件），并覆盖当前进程的镜像、数据和堆栈等信息，包括进程的PID。这用用户的进程会替换掉init进程。
	*/
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}
	return nil
}
