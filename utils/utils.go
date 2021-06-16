package utils

import (
	"cronProject/global"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
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
		global.Logger.Info("Cmd-start")
	}
	result, _ := ioutil.ReadAll(stdout)
	resdata := string(result)
	global.Logger.Info("Cmd" + resdata)
	var res int
	if err := cmd.Wait(); err != nil {
		if ex, ok := err.(*exec.ExitError); ok {
			res = ex.Sys().(syscall.WaitStatus).ExitStatus() //获取命令执行返回状态，相当于shell: echo $?
		}
	}
	global.Logger.Info("Cmd-res" + strconv.Itoa(res))
}
