package controllers

import (
	"github.com/astaxie/beego"
)

type PlaylistController struct {
	beego.Controller
}

func (c *PlaylistController) Get() {
	c.Data["Playlist"] = "DOOD!"
}
