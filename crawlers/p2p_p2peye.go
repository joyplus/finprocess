package crawlers


import (
	"net/http"
	"io/ioutil"
	"fmt"
	"github.com/henrylee2cn/mahonia"
)

func EyeCrawler() {
	u := "http://www.p2peye.com/platform/search/h0i0c0x0r0t0s0b0p1.html"
	resp, _ := http.Get(u)
	body, _ := ioutil.ReadAll(resp.Body)
	s := string(body)
	dec := mahonia.NewDecoder("GBK")
	html := dec.ConvertString(s)

	fmt.Println(html)

	defer resp.Body.Close()

}
