package des1

import "strings"

func simplifyPath(path string) string {
	components := strings.Split(path, "/")
	var res []string
	for i := 0; i < len(components); i++ {
		directory := components[i]
		if directory == "." || directory == "" {
			continue
		}
		if directory == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
			continue
		}
		res = append(res, directory)
	}

	return "/" + strings.Join(res, "/")
}
