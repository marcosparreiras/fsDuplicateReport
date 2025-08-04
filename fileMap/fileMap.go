package fileMap

import "fmt"

type FileMap map[string][]fileInfo

func New() FileMap {
	return make(map[string][]fileInfo)
}

func (fm FileMap) AddFileInfo(name string, path string, size int64) {
	_, ok := fm[name]
	fi := fileInfo{path, size}
	if !ok {
		fm[name] = []fileInfo{fi}
	} else {
		fm[name] = append(fm[name], fi)
	}
}

func (fm FileMap) String() string {
	var output string
	for key, value := range fm {
		if len(value) < 2 {
			continue
		}
		output += fmt.Sprintf("%s:\n", key)
		for i, v := range value {
			output += fmt.Sprintf("\t%d: %v\n", i, v.String())
		}
		output += "\n"
	}
	return output
}
