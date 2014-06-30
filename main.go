package main

import (
	"flag"
	"log"
	"os"

	"github.com/vinaymayar/sgit/cache"
	"github.com/vinaymayar/sgit/config"
	"github.com/vinaymayar/sgit/git"
	"github.com/vinaymayar/sgit/utils"
)

func main() {
	flag.Parse()
	parse(os.Args[1:])
}

func init() {
	utils.MakeSgitRootDir()
	err := utils.NavToGitRootDir()
	if err != nil {
		log.Fatal(err)
	}
}

func parse(args []string) {
	if len(args) > 0 {
		cmd := args[0]
		switch {
		case cache.IsCacheCmd(cmd):
			git.RunWithCache(args)
		case config.IsConfigCmd(cmd):
			config.Configure()
		default:
			git.Run(args)
		}
	} else {
		help()
	}
}

func help() {
}
