package {{.PackageName}}

import "github.com/wetrycode/tegenaria"

type {{.Pipeline}}Pipeline struct {
	Priority int
}

func (p *{{.Pipeline}}Pipeline) ProcessItem(spider tegenaria.SpiderInterface, item *tegenaria.ItemMeta) error {
	return nil

}

func (p *{{.Pipeline}}Pipeline) GetPriority() int {
	return p.Priority
}
