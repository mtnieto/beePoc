package controllers

import (
	"github.com/beePoc/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Series
type SerieController struct {
	beego.Controller
}

// @Title createSerie
// @Description create series
// @Param	body		body 	models.Serie	true		"body for serie content"
// @Success 200 {string} models.Serie.Name
// @Failure 403 body is empty
// @router / [post]
func (u *SerieController) Post() {
	var serie models.Serie
	json.Unmarshal(u.Ctx.Input.RequestBody, &serie)
	uid := models.AddOne(serie)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Series stored
// @Success 200 {object} models.Serie
// @router / [get]
func (u *SerieController) GetAll() {
	series := models.GetAll()
	u.Data["json"] = series
	u.ServeJSON()
}

// @Title Get
// @Description get Serie by uid / in this case  its Name
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Serie
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *SerieController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		serie, err := models.GetOne(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = serie
		}
	}
	u.ServeJSON()
}

// // @Title Update
// // @Description update the user
// // @Param	uid		path 	string	true		"The uid you want to update"
// // @Param	body		body 	models.User	true		"body for user content"
// // @Success 200 {object} models.User
// // @Failure 403 :uid is not int
// // @router /:uid [put]
// func (u *SerieController) Put() {
// 	uid := u.GetString(":uid")
// 	if uid != "" {
// 		var user models.Serie
// 		json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 		uu, err := models.Update(uid, &user)
// 		if err != nil {
// 			u.Data["json"] = err.Error()
// 		} else {
// 			u.Data["json"] = uu
// 		}
// 	}
// 	u.ServeJSON()
// }

// @Title Delete
// @Description delete the serie
// @Param	uid		path 	string	true		"The serie you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *SerieController) Delete() {
	uid := u.GetString(":uid")
	models.Delete(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

