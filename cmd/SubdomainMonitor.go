package main

import (
	"SRCdomian_monitor/pkg/core"
	"fmt"
	time2 "time"
)

func main() {
	time := time2.Now()
	//everything start here
	core.Start()
	elapsed := time2.Since(time)
	fmt.Println("⏰ cost time："+elapsed.String())
}
