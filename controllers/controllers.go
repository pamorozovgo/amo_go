package controllers

import (
	"github.com/astaxie/beego"
	. "../models"
)

type MainController struct {
	beego.Controller
}

func (m *MainController) Index() {
	m.TplName = "templates/layout.html"
}

func (m *MainController) Show(counter *Counter) uint64 {
	return counter.Get()
}

func (m *MainController) Add(counter *Counter) {
	counter.Add()
}