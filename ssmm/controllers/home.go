package controllers

// HomeRouter serves home page.
type HomeController struct {
	baseController
}

// Get implemented Get method for HomeRouter.
func (this *HomeController) Get() {
	this.Data["IsHome"] = true
	this.TplName = "home.html"
}

type DownloadsController struct {
	baseController
}

func (this *DownloadsController) Get() {
	this.Data["IsDownload"] = true
	this.TplName = "downloads.html"
}

type TutorialController struct {
	baseController
}

func (this *TutorialController) Get() {
	this.Data["IsQuickStart"] = true
	this.TplName = "tutorial.html"
}

type TosController struct {
	baseController
}

func (this *TosController) Get() {
	this.TplName = "tos.html"
}

type AboutController struct {
	baseController
}

func (this *AboutController) Get() {
	this.Data["IsAbout"] = true
	this.TplName = "about.html"
}
