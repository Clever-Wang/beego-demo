package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//测试表单验证
type FormValidationController struct {
	beego.Controller
}

type User struct {
	Name string `form:"name"`
	Age int `form:"age"`
	Address string `form:"address"`
}

func (c *FormValidationController) Get() {
	c.TplName = "form.tpl"
}

func (c *FormValidationController) Post() {

	//将表单数据绑定到User
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		log.Println("数据绑定错误")
	}

	valid := validation.Validation{}
	valid.Required(u.Name, "name")
	valid.MaxSize(u.Name, 15, "nameMax")
	valid.Range(u.Age, 0, 18, "age")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	// or use like this
	if v := valid.Max(u.Age, 140, "age"); !v.Ok {
		log.Println(v.Error.Key, v.Error.Message)
	}
	// 定制错误信息
	minAge := 18
	valid.Min(u.Age, minAge, "age").Message("少儿不宜！")
	// 错误信息格式化
	valid.Min(u.Age, minAge, "age").Message("%d不禁", minAge)

	c.StopRun()

}

//通过 StructTag 使用示例

// 验证函数写在 "valid" tag 的标签里
// 各个函数之间用分号 ";" 分隔，分号后面可以有空格
// 参数用括号 "()" 括起来，多个参数之间用逗号 "," 分开，逗号后面可以有空格
// 正则函数(Match)的匹配模式用两斜杠 "/" 括起来
// 各个函数的结果的 key 值为字段名.验证函数名
type user struct {
	Id     int
	Name   string `form:"name" valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Age    int    `form:"age" valid:"Range(1, 140)"` // 1 <= Age <= 140，超出此范围即为不合法
	Email  string `form:"email" valid:"Email; MaxSize(100)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile string `form:"mobile" valid:"Mobile"` // Mobile 必须为正确的手机号
	IP     string `form:"ip" valid:"IP"` // IP 必须为一个正确的 IPv4 地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}

func (c *FormValidationController) ValidStructTag() {
	//将表单数据绑定到User
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		log.Println("数据绑定错误")
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.StopRun()
}





