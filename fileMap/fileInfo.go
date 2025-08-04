package fileMap

import "fmt"

type fileInfo struct {
	path string
	size int64
}

func (fi *fileInfo) String() string {
	return fmt.Sprintf("%s [%.2f Kb]", fi.path, float64(fi.size)/1e3)
}
