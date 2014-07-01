package cache

var gitCacheCmds = map[string]bool{
	"checkout": true,
}

var clearCacheCmd = "clear-cache"

func IsGitCacheCmd(cmd string) bool {
	return gitCacheCmds[cmd]
}

func IsClearCacheCmd(cmd string) bool {
	return cmd == clearCacheCmd
}
