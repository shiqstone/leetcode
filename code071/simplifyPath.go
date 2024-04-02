package code071

import (
	"strings"
)

func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	stack := []string{}

	for _, dir := range dirs {
		switch dir {
		case "", ".":
			// Ignore empty or current directory
			continue
		case "..":
			// Move one directory up if possible
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			// Push valid directory to stack
			stack = append(stack, dir)
		}
	}

	// Construct simplified path
	simplifiedPath := "/" + strings.Join(stack, "/")
	return simplifiedPath
}
