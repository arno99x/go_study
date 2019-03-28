package fetcher

import (
	"bufio"
	"errors"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimit = time.Tick(100 * time.Millisecond)

func Fetch(url string) (content []byte, err error) {
	//爬取数据太快的话，有的网站会对你采取一系列的挫失
	//添加时间隔来限制爬取流量，间隔100毫秒，也就是说如果有10个worker，每10毫秒就会请求10次页面
	<-rateLimit

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = errors.New(err.Error())
		return
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		err = errors.New(err.Error())
		return
	}
	defer resp.Body.Close()

	//网页的编码是未知的，可能是GBK、UTF8等，所以要是能包装个方法去判断网络编码就好了
	//正好golang.org/net/html包中有个chartset.DetermineEncoding方法可以直接拿来使用
	//把现在把这个方法给重新包装成DetermineEncoding()来用
	bodyReader := bufio.NewReader(resp.Body)
	e := DetermineEncoding(bodyReader)

	//fmt.Printf("bodyReader -> \n%s",bodyReader)
	//fmt.Printf("\nresp.body -> \n%s",resp.Body)

	//用对应的编码读取网页编码
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	if resp.StatusCode != http.StatusOK {
		err = errors.New("this is a bad request")
		return
	}
	//fmt.Printf("===============\nutf8Reader ->\n%s",utf8Reader)
	content, err = ioutil.ReadAll(utf8Reader)
	//fmt.Printf("\ncontent -> %s\n",content)
	if err != nil {
		return
	}

	return
	//正则功能：获取网页中所有城市的列表
	//城市列表网页元素样例特征提取 ：
	// 		<a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f="">阿坝</a>
	//		<a href="http://www.zhenai.com/zhenghun/akesu" data-v-5e16505f="">阿克苏</a>
	//		。。。。
	//
	//需求： 提取出所有这样的城市列表
	//正则编写： `<a href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+"[^>]*>[^<]+</a>`
	//正则解读：
	//		所有列表前面长的都一样，`<a href="http://www.zhenai.com/zhenghun/`
	//		匹配0个或多个非>字符并到第一个>结束，`<a href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+"[^>]*>`
	//		匹配至少一个字符并到第一个<的地方结束，最包含</a>, <a href="http://www.zhenai.com/zhenghun/aba" data-v-5e16505f="">阿坝</a>
	//		正则中有两个小括号括起来的部分，分别是：URL->(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)  CITY -> ([^<]+)
	//r := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+)"[^>]*>([^<]+)</a>`)
	//cityList := r.FindAllSubmatch(all,-1)
	//for _,item := range cityList {
	//	fmt.Printf("CITY : %s    ; URL : %s \n",item[2],item[1])
	//}
}

//判断网络编码
func DetermineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(bytes, "")
	return encode
}
