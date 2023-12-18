package main

import (
	"crypto/tls"
	"{{.ModuleName}}/middlewares"
	"{{.ModuleName}}/pipelines"
	"{{.ModuleName}}/spiders"

	"github.com/wetrycode/tegenaria"
)

func main() {
	{{.SpiderName}}SpiderInstance := &spiders.{{.SpiderName}}Spider{
		Name:     "example",
		FeedUrls: []string{"http://www.demo.com/"},
	}
	// 设置下载组件
	Downloader := tegenaria.NewDownloader(tegenaria.DownloadWithTLSConfig(&tls.Config{InsecureSkipVerify: true, MaxVersion: tls.VersionTLS12}))
	Engine := tegenaria.NewEngine(tegenaria.EngineWithDownloader(Downloader))
	Engine.RegisterSpiders({{.SpiderName}}SpiderInstance)
	pipe := &pipelines.{{.SpiderName}}Pipeline{
		Priority: 1,
	}
	Engine.RegisterPipelines(pipe)

	middleware := &middlewares.{{.SpiderName}}DownloadMiddler{
		Priority: 1,
	}
	Engine.RegisterDownloadMiddlewares(middleware)
}
