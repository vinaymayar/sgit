package cache

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vinaymayar/sgit/utils"
)

var targets = map[string]string{
	"sbt": "target",
	"":    "bin",
}

func SaveCache() {
	settings, err := getCacheSettings()
	if err != nil && os.IsNotExist(err) {
		log.Print(err)
		return
	}

	targetDirs, err := getTargetDirs(settings)
	if err != nil {
		log.Print(err)
		return
	}

	for _, dir := range targetDirs {
		if 0 == len(dir) {
			continue
		}

		relDir, err := filepath.Rel(settings.rootDir, dir)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Saving %v.\n", relDir)

		_, err = utils.Execute("rsync", "-Rr", relDir, settings.cachePrefix)
		if err != nil {
			log.Print(err)
		}
	}
}

func RestoreCache() {
	settings, err := getCacheSettings()
	if err != nil {
		log.Print(err)
		return
	}

	err = os.Chdir(settings.cachePrefix)
	if err != nil {
		log.Print(noCacheError(settings.branch))
	}
	cacheDirs, err := getCacheDirs(settings)
	if err != nil {
		log.Print(err)
		return
	}

	for _, dir := range cacheDirs {
		if 0 == len(dir) {
			continue
		}

		relDir, err := filepath.Rel(settings.cachePrefix, dir)
		if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("Restoring %v.\n", relDir)

		_, err = utils.Execute("rsync", "-Rr", relDir, settings.rootDir)
		if err != nil {
			log.Print(err)
		}
	}
}

func ClearCache(args []string) {
	sgitDir, err := utils.GetSgitRootDir()
	if err != nil {
		log.Fatal(err)
	}

	if len(args) == 1 {

		utils.Execute(
			"find",
			sgitDir,
			"-d",
			"-mindepth", "1",
			"-type", "d",
			"-exec", "rm", "-r", "{}", ";")

	} else {

		for _, branch := range args[1:] {
			cachePrefix := filepath.Join(sgitDir, branch)
			_, err = utils.Execute("rm", "-r", cachePrefix)

			if err != nil && os.IsNotExist(err) {
				log.Print(noCacheError(branch))
			}
		}

	}
}

func getCacheDirs(settings CacheSettings) ([]string, error) {
	cacheDirs, err := utils.Execute(
		"find",
		settings.cachePrefix,
		"-name",
		targets[settings.config.ProjectType])
	return strings.Split(strings.TrimSpace(cacheDirs), "\n"), err
}

func getTargetDirs(settings CacheSettings) ([]string, error) {
	targetDirs, err := utils.Execute(
		"find",
		settings.rootDir,
		"-name",
		targets[settings.config.ProjectType],
		"-not",
		"-path",
		filepath.Join(settings.sgitDir, "*"))

	return strings.Split(strings.TrimSpace(targetDirs), "\n"), err
}
