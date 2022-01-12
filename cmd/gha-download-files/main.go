package main

import (
	"github.com/mozartatplay/github-analysis/config"
	"github.com/mozartatplay/github-analysis/githubarchive"
)

func main() {
	cfg := config.Read("config.toml")
	githubarchive.DownloadFiles(cfg.GithubarchivePath)
}
