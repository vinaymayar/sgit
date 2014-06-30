package cache

import "fmt"

func noCacheError(branch string) error {
	return fmt.Errorf("Could not locate cache for %v.\n", branch)
}
