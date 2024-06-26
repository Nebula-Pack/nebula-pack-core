package models

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ScanRockspec(path string) (map[string]string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	dependencies := make(map[string]string)
	lines := strings.Split(string(data), "\n")
	inDependencies := false

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "dependencies") {
			inDependencies = true
			continue
		}
		if inDependencies {
			if line == "}" {
				break
			}
			line = strings.Trim(line, ",")
			line = strings.Trim(line, "\"")
			parts := strings.SplitN(line, " ", 2)
			if len(parts) == 2 {
				dependencies[parts[0]] = strings.TrimSpace(parts[1])
			}
		}
	}

	return dependencies, nil
}
