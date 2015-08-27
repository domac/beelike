package routers

import (
	"beelike/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})

	beego.Router("/", new(controllers.BuoyController), "get:Index")
	beego.Router("/buoy/retrievestation", new(controllers.BuoyController), "post:RetrieveStation")
	beego.Router("/buoy/station/:stationId", new(controllers.BuoyController), "get,post:RetrieveStationJSON")
}
