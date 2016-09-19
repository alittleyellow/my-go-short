package shortlib

import (
	"container/list"
	"errors"
)

type UrlElement struct {
	Original string
	Short string
}

type LRU struct {
	HashShortUrl map[string]*list.Element
	HashOriginUrl map[string]*list.Element
	ListUrl *list.List
	RedisCli *RedisAdaptor
}

func NewLRU(redis_cli *RedisAdaptor) (*LRU, error) {
	lru := &LRU{make(map[string]*list.Element), make(map[string]*list.Element), list.New(), redis_cli}
	return lru, nil
}

func (this *LRU) GetOriginalUrl(short_url string) (string, error) {
	element, ok := this.HashShortUrl[short_url]
	if ok == false {
		original_url, err := this.RedisCli.GetUrl(short_url)
		if err != nil {
			return "", errors.New("No URL")
		}
		//更新LRU缓存
		ele := this.ListUrl.PushFront(UrlElement{original_url, short_url})
		this.HashShortUrl[short_url] = ele;
		this.HashOriginUrl[original_url] = ele;
		return original_url, nil;
	}
	//调整位置
	this.ListUrl.MoveToFront(element)
	ele, _ := element.Value.(UrlElement)

	return ele.Original, nil;
}

func (this *LRU) GetShortUrl(original_url string) (string, error) {
	element, ok := this.HashOriginUrl[original_url]
	if ok == false {
		return "", errors.New("No URL")
	}
	this.ListUrl.MoveToFront(element)
	ele, _ := element.Value.(UrlElement)
	return ele.Short, nil
}

func (this *LRU) SetUrl(original_url, short_url string) error {
	ele := this.ListUrl.PushFront(UrlElement{original_url, short_url})
	this.HashShortUrl[short_url] = ele
	this.HashOriginUrl[original_url] = ele
	err := this.RedisCli.SetUrl(short_url, original_url)
	if err != nil {
		return err
	}
	return nil
}

func (this *LRU) checkList() error {
	return nil
}