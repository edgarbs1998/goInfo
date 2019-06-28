package goInfo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetInfo() *GoInfoObject {
	out := _getInfo()
	osStr := strings.Replace(out,"\n","",-1)
	osStr = strings.Replace(osStr,"\r\n","",-1)
	tmp1 := strings.Index(osStr,"[Version")
	tmp2 := strings.Index(osStr,"]")
	var osInfo [2]string
	osInfo[0] = osStr[0:tmp1-1]
	if tmp1 == -1 || tmp2 == -1 {
		osInfo[1] = "unknown"
	} else {
		osInfo[1] = osStr[tmp1+9:tmp2]
	}
	gio := &GoInfoObject{Kernel:"Windows",Core:osInfo[1],Platform:runtime.GOARCH,OS:osInfo[0],GoOS:runtime.GOOS,CPUs:runtime.NumCPU()}
	gio.Hostname,_ = os.Hostname()
	return gio
}

func _getInfo() string {
	cmd := exec.Command("cmd","ver")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("getInfo:",err)
	}
	return out.String()
}