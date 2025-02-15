package docker

import "strings"

func ExtractImageName(url string) string {
	// extract image name from end-2 & end-3 urls
	// supporting both <org>/<repo> and multi-tier names
	splits := strings.Split(url, "/")
	return strings.Join(splits[1:len(splits)-2], "/")
}
