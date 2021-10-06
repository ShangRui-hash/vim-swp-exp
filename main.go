package main

import (
	"errors"
	"log"
	"net/url"
	"os"
	"vim-swp-exp/settings"
	"vim-swp-exp/watcher"

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
		w := watcher.NewWatcher()
		w.Watch(u)
	} else {
		cli.ShowAppHelp(c)
		return errors.New("pls specify -u or -f")
	}
	return nil
}
