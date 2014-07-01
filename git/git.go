package git

import (
	"log"

	"github.com/vinaymayar/sgit/cache"
	"github.com/vinaymayar/sgit/utils"
)

func Run(args []string) {
	output, errOutput, err := utils.ExecuteWithStderr("git", args...)

	if len(output) > 0 {
		log.Print(output)
	}
	if len(errOutput) > 0 {
		log.Print(errOutput)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func RunWithCache(args []string) {
	cache.SaveCache()
	Run(args)
	cache.RestoreCache()
}
