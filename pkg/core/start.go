package core

import (
	"encoding/json"
)

func Start() {

	ShowBanner()
	//load_yaml
	root_path := GetRootPath()
	Debugf(" 文件绝对路径：%s\n", root_path)
	monitor_yaml := loadConfig(root_path)
	Infof(" Spying domains ：%s\n", monitor_yaml.Monitor.Domain)

	db := Database_init(root_path)
	//获取所有工具跑完的结果，然后放到domainList，最后去重
	//domainList := []string{"wst520.top", "baidu.com", "vip.com"}
	//需用make，不然会数组越界。 https://stackoverflow.com/questions/61189263/panic-runtime-error-index-out-of-range-0-with-length-0
	domainList := make([]DomainList, len(monitor_yaml.Monitor.Domain))
	for k, _ := range monitor_yaml.Monitor.Domain {
		if monitor_yaml.Monitor.Tools.SubdomainBrute_tools.Subfinder_enable == "enable" {
			result_slice, _ := GetSubfinderResult(monitor_yaml.Monitor.Domain[k])
			Debugf("%s\n", result_slice)
			domainList[k].Domain = monitor_yaml.Monitor.Domain[k]
			domainList[k].Subdomain = MergeSlice(domainList[k].Subdomain, result_slice)
		} else {
			Infof("subfinder is disable\n")
		}
		if monitor_yaml.Monitor.Tools.SubdomainBrute_tools.Ksubdomain_enable == "enable" {
			result_slice, _ := GetKsubdomainResult(monitor_yaml.Monitor.Domain[k])
			Debugf("%s\n", result_slice)
			domainList[k].Domain = monitor_yaml.Monitor.Domain[k]
			domainList[k].Subdomain = MergeSlice(domainList[k].Subdomain, result_slice)
			//domainList = MergeSlice(domainList, result_slice)
		} else {
			Infof("ksubdomain_old is disable\n")
		}
	}
	//去重

	for k, _ := range domainList {
		domainList[k].Subdomain = removeDuplicateElement(domainList[k].Subdomain)
		//Debugf(domainList[k].Domain+" subdomain:%s\n",domainList[k].Subdomain)
	}

	//println(len(domainList[0].Subdomain))
	//统计所有子域名个数，如果为0则退出
	var count int
	for v,_ := range domainList{
		count += len(domainList[v].Subdomain)
	}

	if count == 0 {
		Fatalf("found nothing for all domains")
	}

	if monitor_yaml.Monitor.Tools.Web_check_tools.Httpx_enable == "enable" {
		Infof("httpx enable,checking..\n")
		//运行httpx时不注释
		for k, _ := range domainList {
			//逐一域名获取httpx结果,删除工具生成的中间文件
			httpx_result,resultname,filename := GetHttpxResult(domainList[k].Subdomain)
			RemoveFile(resultname)
			RemoveFile(filename)
			//逐条打印json信息
			//for _, v := range httpx_result {
			//	Infof("%s\n", v)
			//}
			httpx_result_json_list := []HttpxResult{}
			for _, v := range httpx_result {
				httpx_result_json := HttpxResult{}
				if v != "" {
					err := json.Unmarshal([]byte(v), &httpx_result_json)
					CheckError(err)
					Debugf(httpx_result_json.URL)
					httpx_result_json_list = append(httpx_result_json_list, httpx_result_json)
				}
			}
			//Debugf("line:%s", httpx_result)
			//获取数据库域名内容
			db_subdomain := GetSubdomainList(db, domainList[k].Domain)
			Debugf("dbsubdomain:%s\n", db_subdomain)
			for k2, _ := range httpx_result_json_list {
				var subdomain Subdomain
				subdomain.Domain = domainList[k].Domain
				subdomain.SubdomainName = httpx_result_json_list[k2].URL
				subdomain.New = 1
				subdomain.Resource = ""
				subdomain.Ip = httpx_result_json_list[k2].IP
				subdomain.Http_status = httpx_result_json_list[k2].StatusCode
				//逐一子域名判断是是否存在于数据库中
				if len(db_subdomain) != 0 {
					if ElementInSlice(httpx_result_json_list[k2].URL, db_subdomain) == true {
						subdomain.New = 0
						UpdateData(db, subdomain)
						continue
					}
				}
				//入库

				SaveData(db, subdomain)

			}
			db_subdomain1 := GetSubdomainList(db, domainList[k].Domain)
			Debugf("[+] now dbsubdomain:%s\n", db_subdomain1)
		}

	}
	if IfNewSubdomainFound(db) {
		new_subdomain_list := SelectNewSubdomain(db)
		for _, v := range new_subdomain_list {
			Debugf("post:%s\n", v)
		}
		Notification(monitor_yaml, new_subdomain_list)
	} else {
		Infof("😢 didn't find any new subdomains\n")
	}
	if monitor_yaml.Monitor.AutoUpadteNewParma == "true"{
		Infof("AutoUpadteNewParma..checking..\n")
		for _,k := range monitor_yaml.Monitor.Domain{
			UpdateAllNewToOld(db,k)
		}
	}
}
