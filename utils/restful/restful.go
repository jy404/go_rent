package restful

const (
	STATUS_CODE_200 = 200
	//STATUS_CODE_400 = 400
	//STATUS_CODE_401 = 401
	//STATUS_CODE_404 = 404
	//STATUS_CODE_500 = 500
)

//func Success(response beego.Controller, data interface{}, msg string) (result map[string]interface{}) {
// todo: 如何传递指定类型数据？
//	result = map[string]interface{}{
//		"code": STATUS_CODE_200,
//		"msg":  msg,
//		"data": data,
//	}
//	response.Data["json"] = result
//	response.ServeJSON()
//	response.StopRun()
//}