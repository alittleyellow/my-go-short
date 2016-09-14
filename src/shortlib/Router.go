package shortlib

import(
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Router struct {
	Configure *Configure
	Processors map[int]Processor
}

const (
	SHORT_URL = 0;
	ORIGINAL_URL = 1;
	UNKOWN_URL = 2;
)

//数据分发
func (this *Router) ServerHttp(w http.ResponseWriter, r *http.Request) {
	start := TimeNow()
	request_url := r.RequestURI[1:]
	action, err := this.ParseUrl(request_url);
	if err != nil {
		fmt.Printf("[ERROR]parse url fail : %v \n", err)
	}
	err = r.ParseForm()
	if err != nil {
		return 
	}
	params := make(map[string]string)
	for k, v := range r.From {
		params[k] = v[0]
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil && err != io.EOF {
		return 
	}
	if r.Method == "GET" {
		action = 0
	} else {
		action = 1
	}

}

func (this *Router) ParseUrl(url string) (action int, err error) {
	if this.isShortUrl(url) {
		return SHORT_URL, nil
	} else {
		return ORIGINAL_URL, nil
	}
}

func (this *Router) isShortUrl(url string) bool {
	short_url_pattern := "XXXX"
	url_reg_exp, err := regexp.Compile(short_url_pattern)
	if err != nil {
		return false
	}
	short_match := url_reg_exp.FindStringSubmatch(url) 
	if short_match == nil {
		return false
	}

	return true
}












