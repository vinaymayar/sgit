package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

const sgitPrefix = "/.git/sgit"

func GetGitRootDir() (string, error) {
	out, err := Execute("git", "rev-parse", "--show-toplevel")
	gitRootDir := strings.TrimSpace(out)

	if err != nil {
		return gitRootDir, err
	}

	if exists, err := pathExists(gitRootDir); exists {
		return gitRootDir, err
	} else if err != nil {
		return gitRootDir, err
	} else {
		err := errors.New("Could not locate git repo.\n")
		return gitRootDir, err
	}
}

func NavToGitRootDir() error {
	gitRootDir, err := GetGitRootDir()
	if err != nil {
		return err
	}
	return os.Chdir(gitRootDir)
}

func MakeSgitRootDir() error {
	sgitRootDir, err := GetSgitRootDir()
	if err != nil {
		return err
	}
	return os.Mkdir(sgitRootDir, 0777)
}

func GetSgitRootDir() (string, error) {
	rootDir, err := GetGitRootDir()
	return filepath.Join(rootDir, sgitPrefix), err
}

func GetBranch() (string, error) {
	branch, err := Execute("git", "rev-parse", "--abbrev-ref", "HEAD")
	return strings.TrimSpace(branch), err
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}
	return !os.IsNotExist(err), err
}
