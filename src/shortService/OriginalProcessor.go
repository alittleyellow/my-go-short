package main 

import(
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"shortlib"
	"fmt"
)

type OriginalProcessor struct {
	*shortlib.BaseProcessor
	Count_Channl chan shortlib.CountChannl
}

const POST string = "POST"
const TOKEN string = "token"
const ORIGINAL_URL string = "original"
const SHORT_URL string = "short"

func (this *OriginalProcessor) ProcessRequest(method, request_url string, params map[string]string, body []byte, w http.ResponseWriter, r *http.Request) error {
	if method != POST {
		return errors.New("create short url must be POST the information")
	}
	original_url, has_original_url := params[ORIGINAL_URL]
	fmt.Println(original_url);
	if !has_original_url {
		return errors.New("Post info errors")
	}

	if !shortlib.IsNormalUrl(original_url) {
		return errors.New("Url is not normal")
	}

	short_url, err := this.createUrl(original_url)

	if err != nil {
		return err
	}
	response, err := this.createResponseJson(short_url)
	if err != nil {
		return err 
	}
	header := w.Header()
	header.Add("Content-Type", "application/json")
	header.Add("charset", "UTF-8")
	io.WriteString(w, response)

	return nil
}

func (this *OriginalProcessor) createUrl(original_url string) (string, error) {
	short, err := this.Lru.GetShortUrl(original_url)
	if err == nil {
		return short, nil
	}
	count, err := this.CountFunction()
	if err != nil {
		return "", err 
	}
	short_url, err := shortlib.TransNumToString(count)
	if err != nil {
		return "", err
	}
	this.Lru.SetUrl(original_url, short_url)
	return short_url, nil
}

func (this *OriginalProcessor) createResponseJson(short_url string) (string, error) {

	json_res := make(map[string]interface{})
	json_res[SHORT_URL] = this.HostName + short_url

	res, err := json.Marshal(json_res)
	if err != nil {
		return "", err
	}

	return string(res), nil
}






