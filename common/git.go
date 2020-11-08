package common

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var Pattern = regexp.MustCompile(`^(?:(?:https:\/\/)?([^:/]+\.[^:/]+)\/|git@([^:/]+)[:/]|([^/]+):)?([^/\s]+)\/([^/\s#]+)(?:((?:\/[^/\s#]+)+))?(?:\/)?(?:#(.+))?`)

var providers = map[string]string{
	"bitbucket": "bitbucket.org",
	"github":    "github.com",
}

func CloneRepo(args []string) (string, error) {
	ssh, _ := getCloneSource(args[0])
	out := "."
	if len(args) > 1 {
		out = args[1]
	}
	c := exec.Command("git", "clone", ssh, out, "--depth", "1")
	c.Stderr = os.Stderr
	err := c.Run()
	return out, err
}

func getCloneSource(src string) (string, string) {
	match := Pattern.FindSubmatch([]byte(src))
	provider := getProvider(match)
	domain := providers[provider]
	user := string(match[4])
	name := string(match[5])
	if domain == "" {
		panic("templar supports GitHub, GitLab, Sourcehut and BitBucket")
	}
	ssh := fmt.Sprintf("git@%s:%s/%s", domain, user, name)
	url := fmt.Sprintf("https://%s/%s/%s", domain, user, name)
	return ssh, url
}

func getProvider(match [][]byte) string {
	indexes := []int{1, 2, 3}
	for _, index := range indexes {
		value := string(match[index])
		if value != "" {
			return value
		}
	}
	return "github"
}
