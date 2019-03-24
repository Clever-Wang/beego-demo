package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"log"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// ===========测试获取请求参数==============
type ParamController struct {
	beego.Controller
}

// 只写一个Get方法 访问 http://localhost:8080/param?request=aaa
func (this *ParamController) Get(){
	request := this.GetString("request")
	this.Ctx.Output.Body([]byte(request))
}

// 将 form 表单数据解析到 struct
// 跳转到 form 表单页面
func (this *ParamController) ToFormPage(){
	this.TplName = "form_struct.tpl"
}

func (this *ParamController) Post(){
	u := user{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
	}
	b, _ := json.Marshal(u)
	this.Ctx.Output.Body(b)
}

// 将request body 解析到 struct
func (this *ParamController) RequestBody(){
	var ob user
	var err error
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		bytes, _ := json.Marshal(ob)
		this.Data["json"] = string(bytes)
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

// 测试文件上传
// 跳转到 form 表单页面
func (this *ParamController) ToUploadPage(){
	this.TplName = "upload_file.tpl"
}

func (this *ParamController) UploadFile() {
	f, h, err := this.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	this.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建

	this.Data["json"] = "{\"msg\":\"success\"}"
	this.ServeJSON()
}

// 从用户请求中直接绑定到指定的对象 请求地址：http://localhost:8080/param/banding?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&student.Name=wangsaichao
func (this *ParamController) Banding() {
	var id int
	this.Ctx.Input.Bind(&id, "id")  //id ==123
	fmt.Println(id)

	var isok bool
	this.Ctx.Input.Bind(&isok, "isok")  //isok ==true
	fmt.Println(isok)

	var ft float64
	this.Ctx.Input.Bind(&ft, "ft")  //ft ==1.2
	fmt.Println(ft)

	ol := make([]int, 0, 2)
	this.Ctx.Input.Bind(&ol, "ol")  //ol ==[1 2]
	fmt.Println(ol)

	ul := make([]string, 0, 2)
	this.Ctx.Input.Bind(&ul, "ul")  //ul ==[str array]
	fmt.Println(ul)

	student := student{}
	this.Ctx.Input.Bind(&student, "student")  //user =={Name:"astaxie"}
	bytes, _ := json.Marshal(student)
	fmt.Println(string(bytes))

	this.Ctx.Output.Body([]byte("helloworld"))
}

type student struct {
	Name string
}