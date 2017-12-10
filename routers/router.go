package routers

import (
	"github.com/astaxie/beego"
	"github.com/gabezeck/playlistr/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getPlaylist", &controllers.PlaylistController{})
}
