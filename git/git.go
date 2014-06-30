package git

import (
	"log"

	"github.com/vinaymayar/sgit/cache"
	"github.com/vinaymayar/sgit/utils"
)

func Run(args []string) {
	_, err := utils.Execute("git", args...)
	if err != nil {
		log.Fatal(err)
	}
}

func RunWithCache(args []string) {
	cache.SaveCache()
	Run(args)
	cache.RestoreCache()
}
