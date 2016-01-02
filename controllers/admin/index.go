package admin

type IndexController struct {
	BaseController
}

func (this *IndexController) Get() {
	this.Display("index.tpl")
}
