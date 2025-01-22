package main

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

func main() {
	repoURL := "https://gitee.com/jianlu8023/nunu"
	repoURLWithGit := "https://gitee.com/jianlu8023/nunu.git"

	// 使用 url.Parse 解析 URL
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	parsedURLWithGit, err := url.Parse(repoURLWithGit)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	// 获取路径部分
	repoPath := parsedURL.Path
	repoPathWithGit := parsedURLWithGit.Path

	// 使用 path.Base 获取最后一个 / 后面的内容
	repoName := path.Base(repoPath)
	repoNameWithGit := path.Base(repoPathWithGit)

	fmt.Println("Repo Name:", repoName)
	fmt.Println("Repo Name with .git:", repoNameWithGit)

	// 如果需要处理带 .git 的情况，可以这样做
	if strings.HasSuffix(repoNameWithGit, ".git") {
		repoNameWithGit = strings.TrimSuffix(repoNameWithGit, ".git")
	}
	fmt.Println("Repo Name with .git (trimmed):", repoNameWithGit)

}
