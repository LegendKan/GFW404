package controllers

import (
	"fmt"
	"ssmm/models"
	"strings"
)

// HomeRouter serves home page.
type ServiceController struct {
	baseController
}

// Get implemented Get method for HomeRouter.
func (this *ServiceController) Get() {
	this.Data["IsService"] = true
	this.Data["Title"] = "产品与服务"
	this.TplName = "service.html"
}

func (c *ServiceController) GetActive() {
	var sortby []string
	var order []string
	var limit int64 = 100
	var offset int64 = 0
	fields := []string{"Id", "Location", "Title", "Description", "Month", "Quarter", "Year", "Amount", "Have"}
	query := map[string]string{
		"IsOnline": "1",
	}
	l, _ := models.GetAllActiveServer(query, fields, sortby, order, offset, limit)
	fmt.Println("hello", len(l))
	for i := 0; i < len(l); i++ {
		//server := s.(models.Server)
		//tmp := server.Description
		l[i].Description = strings.Replace(l[i].Description, "|", "<br /><br />", -1)
		l[i].Have=l[i].Amount-l[i].Have
		//server.Description = tmp
		fmt.Println(l[i].Description)
	}
	//fmt.Println(l[0].Description)
	c.Data["services"] = l
	c.Data["Title"] = "服务与产品"
	c.Data["IsService"] = true
	c.TplName = "service.html"
}
