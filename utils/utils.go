package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"syscall"
)

func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}

func Cmd(cmdStr string) {
	cmd := exec.Command("sh", "-c", cmdStr)
	stdout, _ := cmd.StdoutPipe()
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		fmt.Printf("cmd.Start: %v")
	}
	result, _ := ioutil.ReadAll(stdout)
	resdata := string(result)
	fmt.Println("result=", resdata)
	var res int
	if err := cmd.Wait(); err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			fmt.Println("cmd exit status")
			res = ex.Sys().(syscall.WaitStatus).ExitStatus() //获取命令执行返回状态，相当于shell: echo $?
		}
	}
	fmt.Println(res)
}
