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
