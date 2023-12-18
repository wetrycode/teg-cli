package {{.PackageName}}

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wetrycode/tegenaria"
)

var {{.SpiderName}}Log *logrus.Entry = tegenaria.GetLogger("{{.SpiderName}}")

// {{.SpiderName}}Spider 定义一个spider
type {{.SpiderName}}Spider struct {
	// Name 爬虫名
	Name string
	// 种子urls
	FeedUrls []string
}

// {{.SpiderName}}Spider tegenaria item
type {{.SpiderName}}Item struct {

}
func New{{.SpiderName}}Spider(name string,feedUrls []string) *{{.SpiderName}}Spider {
	return &{{.SpiderName}}Spider{
		Name: name,
		FeedUrls: feedUrls,
	}
}
// StartRequest 爬虫启动，请求种子urls
func (e *{{.SpiderName}}Spider) StartRequest(req chan<- *tegenaria.Context) {
	for _, url := range e.GetFeedUrls() {
		// 生成新的request 对象
		{{.SpiderName}}Log.Infof("request %s", url)
		request := tegenaria.NewRequest(url, tegenaria.GET, e.Parser)
		// 生成新的Context
		ctx := tegenaria.NewContext(request, e)
		// 将context发送到req channel
		time.Sleep(time.Second)
		req <- ctx
	}
}

// Parser 默认的解析函数
func (e *{{.SpiderName}}Spider) Parser(resp *tegenaria.Context, req chan<- *tegenaria.Context) error {
	return nil
}

// ErrorHandler 异常处理函数,用于处理数据抓取过程中出现的错误
func (e *{{.SpiderName}}Spider) ErrorHandler(err *tegenaria.Context, req chan<- *tegenaria.Context) {

}

// GetName 获取爬虫名
func (e *{{.SpiderName}}Spider) GetName() string {
	return e.Name
}

// GetFeedUrls 获取种子urls
func (e *{{.SpiderName}}Spider) GetFeedUrls() []string {
	return e.FeedUrls
}
