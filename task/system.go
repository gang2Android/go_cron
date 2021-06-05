package task

import (
	"cronProject/utils/sms"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"syscall"
	"time"
)

type TaskSystem struct {
	Name      string `json:"name"`
	Spec      string `json:"spec"`
	Path      string `json:"path"`
	Min       int    `json:"min"`
	SmsAk     string `json:"sms_ak"`
	SmsAs     string `json:"sms_as"`
	SmsEnd    string `json:"sms_end"`
	SmsName   string `json:"sms_name"`
	SmsCode   string `json:"sms_code"`
	SmsMobile string `json:"sms_mobile"`
}

func (ts TaskSystem) getSystemDisk() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "start", ts.Name)
	usage := diskUsage(ts.Path)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), "end", ts.Name, "结果=[总："+
		strconv.Itoa(int(usage.All))+"已使用："+
		strconv.Itoa(int(usage.Used))+"空闲："+
		strconv.Itoa(int(usage.Free))+"]")
	if int(usage.Free) <= ts.Min {
		_ = sms.NewSms().SendVerCode(ts.SmsAk, ts.SmsAs, ts.SmsEnd, ts.SmsName, ts.SmsCode, ts.SmsMobile, "666666")
	}
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// DiskUsage disk usage of path/disk
func diskUsage(path string) (disk DiskStatus) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return
	}
	disk.All = (fs.Blocks * uint64(fs.Bsize)) / 1024 / 1024 / 1024
	disk.Free = (fs.Bfree * uint64(fs.Bsize)) / 1024 / 1024 / 1024
	disk.Used = disk.All - disk.Free
	return
}

func loadSysDiskTask(tasks *[]TaskSystem) {
	fileLocker.Lock()
	data, err := ioutil.ReadFile("./task_disk.json")
	fileLocker.Unlock()
	if err != nil {
		fmt.Println("read json file error")
		return
	}
	dataJson := []byte(data)
	err = json.Unmarshal(dataJson, tasks)
	if err != nil {
		fmt.Println("unmarshal json file error")
		return
	}
	return
}
