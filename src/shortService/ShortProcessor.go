package main 

import (
	"fmt"
	"net/http"
	"shortlib"
)

type ShortProcessor struct {
	*shortlib.BaseProcessor
}

func (this *ShortProcessor) ProcessRequest(method, request_url string, params map[string]string, body[]byte, w http.ResponseWriter, r *http.Request) error {
	err := shortlib.IsShortUrl(request_url)
	if err != nil {
		return err
	}
	fmt.Println("short URL");
	original_url, err := this.GetOriginalUrl(request_url) 
	if err != nil {
		return err
	}
	fmt.Printf("REQUEST_URL: %v --- ORIGINAL_URL : %v \n", request_url, original_url)
	http.Redirect(w, r, original_url, http.StatusMovedPermanently)
	return nil
}

func (this *ShortProcessor) GetOriginalUrl(request_url string) (string, error) {
	original_url, err := this.Lru.GetOriginalUrl(request_url);
	if err != nil {
		return "", err
	}

	return original_url, nil
}