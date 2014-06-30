package cache

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/vinaymayar/sgit/config"
	"github.com/vinaymayar/sgit/utils"
)

var cacheCmds = map[string]bool{
	"checkout": true,
}

var targets = map[string]string{
	"sbt": "target",
	"":    "bin",
}

type CacheSettings struct {
	rootDir     string
	sgitDir     string
	branch      string
	cachePrefix string
	config      config.Config
}

func getCacheSettings() (CacheSettings, error) {
	settings := *new(CacheSettings)
	rootDir, err := utils.GetGitRootDir()
	if err != nil {
		return settings, err
	}
	settings.rootDir = rootDir

	sgitDir, err := utils.GetSgitRootDir()
	if err != nil {
		return settings, err
	}
	settings.sgitDir = sgitDir

	branch, err := utils.GetBranch()
	if err != nil {
		return settings, err
	}
	settings.branch = branch

	settings.cachePrefix = filepath.Join(settings.sgitDir, settings.branch)
	err = os.Mkdir(settings.cachePrefix, 0777)

	config, err := config.GetConfig()
	if err != nil {
		return settings, err
	}
	settings.config = config

	return settings, err
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
		log.Print(dir)
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

func IsCacheCmd(cmd string) bool {
	return cacheCmds[cmd]
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
