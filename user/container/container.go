package container

import (
	"MsgGO/user/auth"
	"MsgGO/user/dbops"
	"MsgGO/user/dbops/defsModel"
	"MsgGO/user/defs"
	"MsgGO/user/result"
	"MsgGO/user/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**
* @api {POST} /user/registered 用户注册
* @apiParam {String} name 用户名称
* @apiParam {String} password 密码
* */
func UserRegistered(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	// 用户请求数据获取
	r.ParseForm()
	name := utils.StrFilter(r.PostForm.Get("name"))
	password := utils.StrFilter(r.PostForm.Get("password"))
	if name == "" || password == "" {
		result.RequestError(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	salt, _ := utils.NewUUIDSimplicity()
	password = password + salt
	password = utils.Md5Encode(password)


	user := &defsModel.ImUser{Name: name, Password: password, Salt: salt}
	// 查询是否重复
	b, _ := dbops.Engine.Where("name = ?", name).Get(new(defsModel.ImUser))
	if b == false {
		_, e := dbops.Engine.InsertOne(user)
		if e != nil {
			result.RequestError(w,defs.ErrorDBError)
			return
		}
		result.RequestSuccess(w,defs.SuccessRegisterOK)
		return
	}else {
		result.RequestError(w,defs.ErrorRegister)
	}
}

/**
* @api {POST} /user/userlogin 用户登陆
* @apiParam {String} name 用户名称
* @apiParam {String} password 密码
* */
func UserLogin(w http.ResponseWriter,r *http.Request,p httprouter.Params) {
	// 获取数据
	r.ParseForm()
	name := utils.StrFilter(r.PostForm.Get("name"))
	password := utils.StrFilter(r.PostForm.Get("password"))
	if name=="" || password=="" {
		result.RequestError(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	//获取user data
	user := &defsModel.ImUser{Name: name}
	has, e := dbops.Engine.Get(user)
	if e != nil || has == false {
		result.RequestError(w,defs.ErrorRegister)
		return
	}
	user.Password = utils.Md5Encode(password + user.Salt)
	has, e = dbops.Engine.Get(user)
	if has == true && e== nil {
		// 当用户验证成功 生成jwt
		if s, e := auth.GenerateToken(user);e!= nil{
			result.RequestError(w,defs.ErrorGenerateToken)
			return
		}else{
			// 写入数据库
			user.Token = s
			_, e := dbops.Engine.Where("name = ?", user.Name).Update(user)
			if e != nil {
				result.RequestError(w,defs.ErrorDBError)
				return
			}
			data := &map[string]interface{}{
				"token":s,
				"code":200,
			}

			//response := defs.SuccessResponse{HttpSC:http.StatusOK,Data:defs.Success{s,"200"}}
			//result.RequestSuccess(w,response)
			result.RequestInterFace(w,200,data)
			return
		}
	}
	result.RequestError(w,defs.ErrorRegister)
}
