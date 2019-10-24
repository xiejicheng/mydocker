package main

import (
	"os/exec"
	"syscall"
	"os"
	"log"
)

func main() {
	cmd := exec.Command("sh")     // 被 fock 出的新进程内的初始命令使用 sh
	cmd.SysProcAttr = &syscall.SysProcAttr{ // Go 语言封装了 clone() 函数的调用
		Cloneflags: syscall.CLONE_NEWUTS, // 通过参数创建UTS Namespace
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
