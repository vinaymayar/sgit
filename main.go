package main

import (
	"log"
	"os"

	"github.com/vinaymayar/sgit/cache"
	"github.com/vinaymayar/sgit/config"
	"github.com/vinaymayar/sgit/git"
	"github.com/vinaymayar/sgit/utils"
)

var logFlags = 0

func main() {
	parse(os.Args[1:])
}

func init() {
	log.SetFlags(logFlags)
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
		case cache.IsClearCacheCmd(cmd):
			cache.ClearCache(args)
		case cache.IsGitCacheCmd(cmd):
			git.RunWithCache(args)
		case config.IsConfigCmd(cmd):
			config.Configure(args)
		default:
			git.Run(args)
		}
	} else {
		help()
	}
}

func help() {
	helpMsg := `usage: sgit <command> [<args>]

For a list of git commands run ` + "`sgit help`" + `.  To print this help message run ` + "`sgit`" + `.

The sgit-specific commands are
	configure <project type>
		Set the project type of this repository.  The available options are (sbt).

	clear-cache [<branch name>...]
		Clear cached files for branches (or all branches if none are specified).
`

	log.Print(helpMsg)
}
