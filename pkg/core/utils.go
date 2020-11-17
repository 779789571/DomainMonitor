package core

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		Errorf("%s", err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func CheckFileExisted(file_path string) bool {
	_, err := os.Stat(file_path)
	if err != nil {
		return false
	}
	return true
}
func CheckError(err error) {
	if err != nil {
		Errorf("something wrong :%s\n", err)
	}
}
func GetTimeNow() string {
	t := time.Now()
	time := t.Format("2006.Jan.02 15:04")
	return time
}
func CheckFileLength(file_path string) int64 {
	fi, err := os.Stat(file_path)
	if err != nil {
		CheckError(err)
		return 0
	}
	Debugf("size:%d", fi.Size())
	return fi.Size()
}
func FileLoad(file_path string, bufSize int) string {
	var res string
	var chunk []byte
	file, err := os.OpenFile(file_path, os.O_RDONLY, 0666)
	CheckError(err)
	defer file.Close()
	r := bufio.NewReader(file)
	b := make([]byte, bufSize)
	for {
		n, err := r.Read(b)
		chunk = append(chunk, b[:n]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				Errorf("%s", err)
				os.Exit(-1)
			}
		}
		//res = string(b)
	}
	res = string(chunk)
	if res != "" {
		return res
	}
	return ""
}
func ExecCmd(command string) (bool, error) {
	//cmd := exec.Command("/Users/jhin/go/src/src_monitor/subBruteTools/subfinder","-d","xiami.com","-silent")
	//测试执行生成域名文件
	//cmd := exec.Command("/Users/jhin/go_learning_project/domain_monitor/subBruteTools/subfinder","-d","kjaskldjwst520111aa.com","-o","kjaskldjwst520111aa.com","-oI","-oJ","-nW")
	//快速
	//cmd := exec.Command("/Users/jhin/go_learning_project/domain_monitor/subBruteTools/subfinder", "-h")
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	Debugf("excuting:%s\n", cmd)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = &out
	//cmd.Dir = dir
	err := cmd.Start()
	if err != nil {
		Errorf("wrong:%s\n", err)
		return false, err
	}
	err = cmd.Wait()
	//out,err := exec.Command(dir,commandName).CombinedOutput()

	return true, err
}
//合并数组
func MergeSlice(s1 []string, s2 []string) []string {
	for _, value := range s2 {
		s1 = append(s1, value)
	}
	return s1
}
//参考http://www.36nu.com/post/329 和 https://www.yuque.com/fz420/golang/ky17s2
func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func RemoveFile(filename string) {
	err := os.Remove(filename)
	Debugf("remove file %s",filename)
	if err != nil {
		Errorf("%s",err)
		//fmt.Println(err)
	}
}

//元素在切片中
func ElementInSlice(ele string,slice []string) bool {
	for _,v := range slice{
		if ele == v{
			return true
		}
	}
	return false
}
