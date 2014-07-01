package cache

import (
	"os"
	"path/filepath"

	"github.com/vinaymayar/sgit/config"
	"github.com/vinaymayar/sgit/utils"
)

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
