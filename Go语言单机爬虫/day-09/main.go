package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func tryGet()  {
	resp,err := http.Get("http://www.baidu.com/")
	if err != nil {
		// 处理异常
	}
	defer resp.Body.Close()  // 函数结束时关闭Body
	body, err := ioutil.ReadAll(resp.Body)  // 读取Body
	fmt.Printf("%s",body)
}

func getWebContent(url string) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9") // 增加请求报文头
	request.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36")
	if err != nil {
		fmt.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	getMovies(body)
}

func getMovies(body []byte) {

	 Name := regexp.MustCompile("<span class=\"title\">([^&]+)</span>")
	 matches := Name.FindAllSubmatch(body,-1)
	for _,m := range matches{
		fmt.Printf("%s\n",m[1])
	}
	fmt.Println(len(matches))

}

func main()  {

	base_url := "https://movie.douban.com/top250?start=%d&filter="
	for i := 0; i < 10; i++ {
		url := fmt.Sprintf(base_url, i*25)
		getWebContent(url)
	}


}
