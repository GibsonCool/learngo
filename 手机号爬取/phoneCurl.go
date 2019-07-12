package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	//filePath := currentFile() + "phoneList/"
	//fileName := filePath + "slice_Uq6_1.txt"

	//readPhoneList("phoneList")

	engine := ConcurrentEngine{
		Scheduler:   SimpleScheduler{},
		WorkerCount: 10,
	}

	//engine.Run("test")
	engine.Run("phoneList")

	//phoneInfo := worker("17121192629")
	//fmt.Println(phoneInfo)
}

func writeFileErrorHanding(fileName string, info string) {

	//加了 O_EXCL 如果文件已经存在会报错
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		//panic(err)
		//fmt.Println(" ile already exit")
		/*
			    os.OpenFile() 注释中最后一句是： If there is an error, it will be of type *PathError.
				所以我们可以对错误异常进行更加详细的处理
		*/
		//Type assertion 类型断言
		if pathError, ok := err.(*os.PathError); !ok {
			//如果类型转换不是 pathError 就直接输出
			panic(err)
		} else {
			fmt.Printf("^%s, %s, %s\n",
				pathError.Op,
				pathError.Err,
				pathError.Path)
		}
		return
	}
	defer file.Close()
	n, err := io.WriteString(file, info)
	if err != nil {
		fmt.Println("n:", n, "err:", err)
	}

}

/*
	获取当前执行文件的目录
*/
func currentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(" Can not get current file info")
	}
	lastIndex := strings.LastIndex(file, "/") + 1
	file = file[:lastIndex]
	return file
}

/*====================== engine ==============================*/
type ConcurrentEngine struct {
	Scheduler   SimpleScheduler
	WorkerCount int
}

func (e *ConcurrentEngine) Run(dirName string) {

	in := make(chan string)          //phone number
	out := make(chan ParsePhoneInfo) // phone info
	e.Scheduler.ConfigureWorkerChan(in)

	registedFilePath := currentFile() + "registed.txt"
	//unRegistedFilePath := currentFile() + "unregisted.txt"

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	//将任务源 request 送入调度器中
	e.Scheduler.readPhoneList(dirName)

	itemNumber := 0
	for {
		parsePhoneInfo, ok := <-out
		if !ok {
			break
		}

		fmt.Println("itemNumber:", itemNumber)
		itemNumber++
		//var resultFileName string
		if parsePhoneInfo.IsRegisted {
			//resultFileName = registedFilePath
			go writeFileErrorHanding(registedFilePath, "itemNumber:"+strconv.Itoa(itemNumber)+"   phoneNumber:"+parsePhoneInfo.phoneInfo)
		} else {
			//resultFileName = unRegistedFilePath
			fmt.Println(parsePhoneInfo.IsRegisted, parsePhoneInfo.phoneInfo)
		}

	}

}

/*====================== worker ==============================*/
func createWorker(in chan string, out chan ParsePhoneInfo) {
	go func() {
		for {
			phoneNumber := <-in
			var parsePhoneInfo ParsePhoneInfo

			parsePhoneInfo = worker(phoneNumber)

			out <- parsePhoneInfo
		}
	}()
}

func worker(phoneNumber string) ParsePhoneInfo {

	targetUrl := "http://www.mdm998.net:8080/api/code/sms?target=2&phone=" + phoneNumber + "&quick_fail=1&p=%5Baid:58FF8D05-D7AA-406F-BA92-1774B51F6AF2%5D,%5Bcode:CN%5D,%5Blan:zh%5D,%5Bsvc:12.3.1%5D,%5Bsvn:iOS%5D,%5Bcvn:5.0.2%5D,%5Bcvc:2058%5D,%5Bchn:ios_cfqp_axyl%5D,[pkg:com.ioscfgame.axyl]&gv=5.0.2&qt_utc=1562558805251"

	// 初始化 proxy http client
	client := AbuyunProxy{AppID: proxyUser, AppSecret: proxyPass}.ProxyClient()

	request, _ := http.NewRequest(http.MethodGet, targetUrl, bytes.NewBuffer([]byte(``)))

	// 设置IP切换头 (只支持 HTTP)
	headers := request.Header
	headers.Set("Proxy-Switch-Ip", "yes")

	response, e := client.Do(request)

	if e != nil {
		return ParsePhoneInfo{false, phoneNumber + "  failed to connect:   " + e.Error() + "\n"}
	}

	bodyByte, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return ParsePhoneInfo{false, phoneNumber + "  读取body的时候发生错误：" + e.Error() + "\n"}
	}

	defer response.Body.Close()

	var result Result
	err := json.Unmarshal(bodyByte, &result)
	if err != nil {
		return ParsePhoneInfo{false, phoneNumber + "  数据转换的时候出错:" + response.Status + "   error:" + err.Error() + "\n"}
	}

	//fmt.Println(result, phoneNumber)
	if result.Status == 0 {
		return ParsePhoneInfo{true, phoneNumber + "\n"}
	} else {

		return ParsePhoneInfo{false, phoneNumber + "  msg：" + result.Msg + "  status:" + strconv.Itoa(result.Status) + "\n"}
	}
}

/*====================== 代理信息 ==============================*/
// 代理服务器
const proxyServer = "http-pro.abuyun.com:9010"

// 代理隧道验证信息
const proxyUser = "H35Y953H5399P04P"
const proxyPass = "06CB6DEBC6AD95FA"

type AbuyunProxy struct {
	AppID     string
	AppSecret string
}

func (p AbuyunProxy) ProxyClient() http.Client {
	proxyUrl, _ := url.Parse("http://" + p.AppID + ":" + p.AppSecret + "@" + proxyServer)
	return http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}

/*====================== scheduler ==============================*/
type SimpleScheduler struct {
	WorkerChan chan string
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan string) {
	s.WorkerChan = c
}

func (s *SimpleScheduler) Submit(phoneNmber string) {
	go func() { s.WorkerChan <- phoneNmber }()
}

func (s *SimpleScheduler) readPhoneList(dirName string) {
	go func() {
		dirPath := currentFile() + dirName
		PthSep := string(os.PathSeparator)

		dir, e := ioutil.ReadDir(dirPath)
		if e != nil {
			panic("目录不存在")
		}

		for _, fi := range dir {
			if fi.IsDir() {
				fmt.Println(fi.Name() + " 是一个文件夹")
			} else {
				fileName := dirPath + PthSep + fi.Name()

				file, err := os.Open(fileName)
				if err != nil {
					panic(err)
				}
				scanner := bufio.NewScanner(file)

				//省略初始条件,递增条件实现while
				for scanner.Scan() {
					s.Submit(scanner.Text())
				}
				fmt.Println(fileName)
			}
		}
		fmt.Println("文件内容读取完毕=========================")
	}()
}

/*====================== model ==============================*/
type ParsePhoneInfo struct {
	IsRegisted bool
	phoneInfo  string
}

type Result struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
