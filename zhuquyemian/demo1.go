package zhuquyemian

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
    @date: 2022/11/19
**/
func Fech(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// 设置请求header和cookie
	// 解决防爬虫 1.用户代理  2.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36")
	req.Header.Set("Cookie", "_ga=GA1.2.497622810.1658154826; __gads=ID=054a197529a927b1:T=1658154828:S=ALNI_MYw98poBtgSobNbNZsCKwqvYdeFQw; .AspNetCore.Antiforgery.b8-pDmTq1XM=CfDJ8NfDHj8mnYFAmPyhfXwJojec2zu4sR8k2BMjGmi9X86GWMc8ItbfOPrpqh3TRXjsoaZI_qYPHSLOQS4gVShVfzI_6zhhU-5OiFdLTeDOr-xaz6QYIRMZDTyTlM7mesdDNDb49Da5Yhbvs5JuQxuc7Qg; Hm_lvt_866c9be12d4a814454792b1fd0fed295=1667638049,1667664830,1668863125; _gid=GA1.2.2065310825.1668863125; __gpi=UID=000007d6d7223f6d:T=1658154828:RT=1668863124:S=ALNI_MaNjWTdNaXApPSyZ4rqN5vzo6viaQ; Hm_lpvt_866c9be12d4a814454792b1fd0fed295=1668873196")
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get req err", err)
		return ""
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http status code", resp.StatusCode)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read err", err)
		return ""
	}
	return string(body)
}

func main() {

}
