package {{.PackageName}}

import (

	"github.com/wetrycode/tegenaria"
)

type {{.Middlerware}}DownloadMiddler struct {
	Priority int

	Name string
}

func (m {{.Middlerware}}DownloadMiddler) GetPriority() int {
	return m.Priority
}

func (m {{.Middlerware}}DownloadMiddler) ProcessRequest(ctx *tegenaria.Context) error {

	return nil
}

func (m {{.Middlerware}}DownloadMiddler) ProcessResponse(ctx *tegenaria.Context, req chan<- *tegenaria.Context) error {

	return nil

}
func (m {{.Middlerware}}DownloadMiddler) GetName() string {
	return m.Name
}
