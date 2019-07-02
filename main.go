package main

import (
	"flag"
	"os"
	"os/exec"
	"time"
)

var typeFlag string
var selectFlag string

func init() {
	flag.StringVar(&selectFlag, "s", "", "当t=pod时,该参数可选为选择器 typename=xxx")
	flag.StringVar(&typeFlag, "t", "node", "监控数据类型（node/pod）")
}

func main() {
	flag.Parse()
	var ch chan int
	//定时任务
	ticker := time.NewTicker(time.Second * 1)
	typeStr := typeFlag
	selector := selectFlag
	go func(typeStr, selector string) {
		for range ticker.C {
			metrics(typeStr, selector)
		}
		ch <- 1
	}(typeStr, selector)
	<-ch
}

func metrics(typeStr, selector string) {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
	switch typeStr {
	case "node":
		cmd = exec.Command("kubectl", "top", "nodes")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	case "pod":
		selectorStr := ""
		if selector != "" {
			selectorStr = "--selector=" + selector
		}
		cmd = exec.Command("kubectl", "top", "pod", selectorStr)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}
}
