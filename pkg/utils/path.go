package utils

import (
	"log"
	"path"
	"path/filepath"
)

// LocalAbsToRoot converts an absolute path to a relative path to the root
func LocalAbsToRoot(abs string, root string) string {

	rootdir, _ := path.Split(root)
	result := abs[len(rootdir):]
	return "/" + result
}

// LocalRelToRoot converts a relative path to a relative path to the root
func LocalRelToRoot(rel string, root string) string {

	abs, err := filepath.Abs(rel)
	if err != nil {
		log.Panic(err)

	}
	return LocalAbsToRoot(abs, root)

}
