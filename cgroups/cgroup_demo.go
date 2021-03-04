package main

import (
	"os/exec"
	"path"
	"os"
	"fmt"
	"io/ioutil"
	"syscall"
	"strconv"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main(){
	//fmt.Println(os.Args[0] == "/proc/self/exe")
	if os.Args[0] == "/proc/self/exe" {
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")//再次执行本程序，所以会重新调用main函数一次，这样就会进入上面的代码段。上面的代码段会阻塞执行，出错后会退出，所以不用担心重复执行下面的代码
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {//cmd.Start 与 cmd.Wait 必须一起使用, cmd.Start 不用等命令执行完成，就结束
		fmt.Println("ERROR", err)
		os.Exit(1)
	} else {
		fmt.Println("fork进程映射在外部命名空间的PID：%v", cmd.Process.Pid)
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit_buildingdocker"), 0755)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit_buildingdocker", "task"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit_buildingdocker", "memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait() //cmd.Wait 等待命令结束
}