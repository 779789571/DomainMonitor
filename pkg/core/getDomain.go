package core

import (
	"strings"
)

func GetSubfinderResult(monitor_domain string) ([]string, error) {
	var subfinderResiltList = []string{}
	dir := GetRootPath() + "/subBruteTools/"
	Debugf("path:%s\n", dir)
	cmd := dir + "subfinder " + "-d " + monitor_domain + " -o " + "subfinder_" + monitor_domain + " -silent"
	result, err := ExecCmd(cmd)
	CheckError(err)
	file_path := GetRootPath() + "/subfinder_" + monitor_domain
	if result == true {
		if CheckFileLength(file_path) == 0 {
			Infof("subfinder found nothing about %s\n", monitor_domain)
		} else {
			Infof("opening subfinder result for %s\n", monitor_domain)
			res := FileLoad(file_path, 1024)
			subfinderResiltList = strings.Split(res, "\n")
			//println(subfinderResiltList)
		}
		//subfinder json format
		//cmd := dir+"subfinder "+"-d "+monitor_domain+" -o "+monitor_domain+" -oJ -nW -silent"
		////result 判断命令执行正确与否
		//Debugf("doing.."+cmd+"\n")
		//result, err := ExecCmd(cmd)
		//CheckError(err)
		//file_path := GetRootPath() + "/"+monitor_domain
		//if result == true {
		//	if CheckFileLength(file_path) == 0 {
		//		Infof("found nothing about %s\n", monitor_domain)
		//	} else {
		//		Infof("opening result for %s\n", monitor_domain)
		//		res := FileLoad(file_path, 1024)
		//		file_list := []string{}
		//		file_list = strings.Split(res, "\n")
		//		//println(list)
		//
		//		for _, v := range file_list {
		//			//println(v)
		//			if v == "" {
		//				continue
		//			}
		//			sub := SubfinderResult{}
		//			err := json.Unmarshal([]byte(v), &sub)
		//			CheckError(err)
		//			Debugf("test:%s\n", sub.Host)
		//
		//			subfinderResiltList = append(subfinderResiltList,sub)
		//		}
		//		println(subfinderResiltList[0].Host)
		//		return subfinderResiltList,err
		//buf, err := ioutil.ReadFile(file_path)
		//}
	}

	return subfinderResiltList, err

}
func GetKsubdomainResult(monitor_domain string) ([]string, error) {
	//todo add ksubdomain
	ksubdomainResult := []string{}
	dir := GetRootPath() + "/subBruteTools/"
	Debugf("path:%s\n", dir)
	cmd := dir + "ksubdomain " + "-d " + monitor_domain + " -o " + "ksubdomain_" + monitor_domain + " -silent"
	result, err := ExecCmd(cmd)
	CheckError(err)
	file_path := GetRootPath() + "/ksubdomain_" + monitor_domain
	if result == true {
		if CheckFileLength(file_path) == 0 {
			Infof("found nothing about %s\n", monitor_domain)
		} else {
			Infof("opening ksubdomain result for %s\n", monitor_domain)
			res := FileLoad(file_path, 1024)
			ksubdomainResult = strings.Split(res, "\n")
			//println(ksubdomainResult)
		}
	}
	return ksubdomainResult, nil
}
