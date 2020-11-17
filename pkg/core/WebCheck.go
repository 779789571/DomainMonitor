package core

import (
	"github.com/go-cmd/cmd"
	"os"
	"strings"
	"time"
)

func GetHttpxResult(domainList []string) ([]string,string,string) {
	root_path := GetRootPath()
	//cmd :="echo "+host+" | "+ root_path +"/subBruteTools/httpx -json -unsafe "
	//创建一个存储域名文件给httpx读取，然后删除
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	name := (tm.Format("2006_01_02_03_04_05_PM"))
	file_name := name + ".txt"
	file, err := os.Create(file_name)
	if err != nil {
		CheckError(err)
	}

	if len(domainList) != 0 {
		for _, host := range domainList {
			host = strings.Replace(host, " ", "", -1)
			host = strings.Replace(host, "\n", "", -1)
			host = strings.Replace(host, "\r", "", -1)
			host = host + "\r\n"
			file.WriteString(host)
		}
	}
	result_name := "httpxResult_"+file_name
	command :=  root_path+"/subBruteTools/httpx -l "+file_name+" -json -o "+result_name
	Debugf("executing %s\n", command)
	c := cmd.NewCmd("bash", "-c", command)
	<-c.Start()
	//切片转字符串后传入字符串类型通道
	//Debugf("get return:%s\n", s.Stdout)

	//json_result,err := ExecCmdWithReturn(cmd)
	//CheckError(err)
	res := FileLoad(result_name,1000)
	defer func() {
		file.Close()
		//自动删除httpx结果
		//RemoveFile(result_name)
		//RemoveFile(file_name)
	}()
	return strings.Split(res,"\n"),result_name,file_name
}

