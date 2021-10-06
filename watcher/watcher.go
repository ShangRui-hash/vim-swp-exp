package watcher

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/asmcos/requests"
)

//Watcher 监控器
type Watcher struct {
	wg sync.WaitGroup
}

//NewWatcher 构造函数
func NewWatcher() *Watcher {
	var wg sync.WaitGroup
	return &Watcher{
		wg: wg,
	}
}

//Watch 监控
func (w *Watcher) Watch(targetURL *url.URL) {
	w.wg.Add(1)
	u := targetURL.Scheme + "://" + targetURL.Host + targetURL.Path + ".swp"
	log.Printf("watching %s\n", u)
	go func() {
		for range time.Tick(time.Millisecond * 500) {
			h := requests.Header{
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
			}
			resp, err := requests.Get(u, h)
			if err != nil {
				continue
			}
			if resp.R.StatusCode == 404 {
				continue
			}
			go w.download(targetURL.Host, targetURL.RawPath, resp.Text())
		}
	}()
	w.wg.Wait()
}

func (w *Watcher) download(domain, rawPath, respText string) {
	fmt.Println("rawPath:", rawPath)
}
