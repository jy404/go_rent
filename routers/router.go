package routers

import (
	"github.com/astaxie/beego"
	"rent_backend/controllers/account"
	"rent_backend/controllers/house"
	"rent_backend/middleware"
)

func init() {
	middleware.CheckLogin()
	beego.Router("/api/account/login", &account.Controller{}, "Post:Login")
	beego.Router("/api/account/user_info", &account.Controller{}, "Post:BindUserInfo")
	beego.Router("/api/account/bind_phone", &account.Controller{}, "Post:BindPhone")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Get:UserInfo")
	beego.Router("/api/account/edit_info", &account.Controller{}, "Post:EditUserInfo")
	beego.Router("/api/account/operation", &account.Controller{}, "Post:Operation")
	beego.Router("/api/account/get_collects", &account.Controller{}, "Get:Collects")
	beego.Router("/api/account/collects_delete", &account.Controller{}, "Post:OperationDelete")
	// house
	beego.Router("/api/house/city_conf", &house.Controller{}, "Get:CityListConf")
	beego.Router("/api/house/index", &house.Controller{}, "Get:HouseIndex")
	beego.Router("/api/house/selects", &house.Controller{}, "Get:Selects")
	beego.Router("/api/house/banners", &house.Controller{}, "Get:BannerList")
	beego.Router("/api/house/search", &house.Controller{}, "Get:SearchHouse")
	beego.Router("/api/house/nearby_houses", &house.Controller{}, "Get:NearbyHouses")
	beego.Router("/api/house/detail/:house_id([0-9]+)", &house.Controller{}, "Get:HouseDetail")
}
