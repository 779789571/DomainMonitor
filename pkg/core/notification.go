package core

import (
	notification "SRCdomian_monitor/pkg/core/bot"
	"os"
	"strconv"
	"strings"
)

func Notification(monitor_yaml Monitor_yaml,newDomainList []string) {
	if Record(newDomainList){
		Infof("result saved in send_log.txt..")
	}
	Debugf("sending...")
	if monitor_yaml.Monitor.ServerJiang.Enable == "enable"{
		Debugf("using server jiang")
		newDomainList_count := len(newDomainList)
		title := "今日份子域名，请食用qvq"
		content := "#### 发现子域名"+strconv.Itoa(newDomainList_count)+"个，请大佬过目～  \r\n #### 详情如下： \r\n  " +
			"\r\n | subdomain | domain | http_status |  \r\n "
		for _,v := range newDomainList{
			var markdown string
			subdomain_slice := strings.Split(v,"|")
			if subdomain_slice[2] == "200"{
				bold_http_status := " **200** "
				markdown = "| "+subdomain_slice[0]+" | "+subdomain_slice[1]+" | "+bold_http_status+"|  \r\n"
			}else {
				markdown = "| "+subdomain_slice[0]+" | "+subdomain_slice[1]+" | "+subdomain_slice[2]+" |  \r\n"
			}
			content = content + markdown

		}
		println(content)
		if notification.PostToServerJiang(title,content,monitor_yaml.Monitor.ServerJiang.ServerJiangApi){
			Infof("😄 post to serverjiang success!\n")
		}
	}
}

func Record(newDomainList []string) bool {
	newdomain_log_file := GetRootPath()+"/send_log.txt"
	if CheckFileExisted(newdomain_log_file) == true{
		log,err := os.OpenFile(newdomain_log_file,os.O_APPEND,0666)
		defer log.Close()
		CheckError(err)
		for _,v := range newDomainList{
			log.WriteString(v+"\n")
		}
	}else {
		log, err := os.Create(newdomain_log_file)
		defer log.Close()
		CheckError(err)
		for _,v := range newDomainList{
			log.WriteString(v+"\n")
		}

	}
	return true
}
