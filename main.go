package main

import (
	"bufio"
	"errors"
	"log"
	"net/url"
	"os"
	"strings"
	"vim-swp-exp/settings"
	"vim-swp-exp/watcher"

	mapset "github.com/deckarep/golang-set"
	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "vim-swp-exp",
		Usage:     "vim swp file exploit",
		UsageText: "vim-swp-exp -u url  | vim-swp-exp -f file",
		Version:   "v0.1",
		Authors: []*cli.Author{{
			Name:  "无在无不在",
			Email: "2227627947@qq.com",
		}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "file",
				Aliases:     []string{"f"},
				Usage:       "specify the url list to watch",
				Destination: &settings.AppConfig.InputFilePath,
			},
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"u"},
				Usage:       "specify the url to watch",
				Destination: &settings.AppConfig.URL,
			},
		},
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Println("app.Run failed,err:", err)
		return
	}
}

func run(c *cli.Context) error {
	log.Println("start watching:")
	if len(settings.AppConfig.URL) > 0 {
		u, err := url.Parse(settings.AppConfig.URL)
		if err != nil {
			log.Println("url.Parse failed,err:", err)
			return err
		}
		watcher.Watch(u)
	} else if len(settings.AppConfig.InputFilePath) > 0 {
		file, err := os.Open(settings.AppConfig.InputFilePath)
		if err != nil {
			log.Println("os.Open failed,err:", err)
			return err
		}
		URLSet := mapset.NewSet()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 {
				continue
			}
			u, err := url.Parse(line)
			if err != nil {
				log.Println("url.Parse failed,err:", err)
				continue
			}
			if URLSet.Contains(u) {
				continue
			}
			URLSet.Add(u)
			go watcher.Watch(u)
		}
		select {}
	} else {
		cli.ShowAppHelp(c)
		return errors.New("pls specify -u or -f")
	}
	return nil
}
