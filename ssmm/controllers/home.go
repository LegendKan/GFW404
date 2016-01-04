package controllers

// HomeRouter serves home page.
type HomeController struct {
	baseController
}

// Get implemented Get method for HomeRouter.
func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplNames = "home.html"
}

type DownloadsController struct {
	baseController
}

func (this *DownloadsController) Get() {
	this.Data["IsDownload"] = true
	this.TplNames = "downloads.html"
}

type TutorialController struct {
	baseController
}

func (this *TutorialController) Get() {
	this.Data["IsQuickStart"] = true
	this.TplNames = "tutorial.html"
}

type TosController struct {
	baseController
}

func (this *TosController) Get() {
	this.TplNames = "tos.html"
}
