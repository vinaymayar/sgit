package config

var configCmd = "configure"

func IsConfigCmd(cmd string) bool {
	return cmd == configCmd
}
