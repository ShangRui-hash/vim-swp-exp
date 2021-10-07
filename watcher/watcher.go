package watcher

import (
	"crypto/md5"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/asmcos/requests"
)

//Watch 监控
func Watch(targetURL *url.URL) {
	temp := strings.Split(targetURL.Path, "/")
	filename := temp[len(temp)-1]
	swpFilename := "." + filename + ".swp"
	path := strings.ReplaceAll(targetURL.Path, filename, swpFilename)
	u := targetURL.Scheme + "://" + targetURL.Host + path
	log.Printf("watching %s\n", u)
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
		go download(targetURL.String(), resp.Text())
	}
}

func download(targetURL, respText string) {
	temp := strings.Split(targetURL, "/")
	dirPath := strings.Join(temp[2:len(temp)-1], "/")
	filename := "." + temp[len(temp)-1] + ".swp"
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Println("os.MkdirAll failed,err:", err)
		return
	}
	filepath := path.Join(dirPath, filename)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if err := ioutil.WriteFile(filepath, []byte(respText), os.ModePerm); err != nil {
			log.Println("ioutil.WriteFile failed,err:", err)
			return
		}
		log.Println("download file success:", filepath)
	} else {
		//文件存在，比对md5哈希值
		buf, err := ioutil.ReadFile(filepath)
		if err != nil {
			log.Println("ioutil.ReadFile failed,err:", err)
			return
		}
		oldfilemd5 := md5.Sum(buf)
		newfilemd5 := md5.Sum([]byte(respText))
		if oldfilemd5 != newfilemd5 {
			if err := ioutil.WriteFile(filepath, []byte(respText), os.ModePerm); err != nil {
				log.Println("ioutil.WriteFile failed,err:", err)
				return
			}
			log.Printf("%s changed\n", filepath)
		}
	}
}
