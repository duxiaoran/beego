package admin

import (
  "myapp/models"
  "time"
  "fmt"
  "myapp/util"
)

type RoleController struct {
	baseController
}

func (this *RoleController) Index()  {
	var (
		page       int
		pagesize   int = 8
		offset     int
		role_list       []*models.Role
	)
	if page, _ = this.GetInt("page"); page < 1 {
		page = 1
	}
	offset = (page - 1) * pagesize
	query := this.o.QueryTable( new(models.Role).TableName())
	count, _ := query.Count()
	query.OrderBy("created").Limit(pagesize,offset).All(&role_list)
	this.Data["role_list"] = role_list
	this.Data["pagebar"] = util.NewPager(page, int(count), pagesize,fmt.Sprintf("/role/index.html"), true).ToString()
	this.Layout = "admin/layouts/starter.html"
	this.TplName = "admin/role/index.html"
}

func (this *RoleController) Save() {
	name := this.GetString("name")
	desc := this.GetString("desc")
	id,_ := this.GetInt("id")
	role := models.Role{}
	role.Name = name
	role.Desc = desc

	if id == 0 {
		role.Created = time.Now()
		if _, err := this.o.Insert(&role);err == nil {
			this.jsonResult(200,"添加成功",nil)
		}else{
			this.jsonResult(201,"添加失败",nil)
		}
	}else{
		role.Id = id
		if _,err := this.o.Update(&role);err == nil {
			this.jsonResult(200,"修改成功",nil)
		}else{
			this.jsonResult(201,"修改失败",nil)
		}
	}

	this.jsonResult(200,name,nil)
}


