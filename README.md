# hiweb
## 编译命令hi
```
	cd hiweb\webcmd\hi
	go build .
```
## 使用命令hi
```
	生成项目     hi init [项目名称]
	生成swagger     hi swag [controllers文件夹] [项目名称]
	生成数据库model  hi gen [mysql链接] [生成model路径]
```
## 一个简单的web框架
```
可以使用注释的方式暴露http接口

type Conference struct {
	hiweb.Controller
}

// @Description create
// @Param name 会议名称
// @Param username 用户名称
// @Param startTime 开始时间
// @Param endTime 结束时间
// @Param password 密码
// @Auth
func (c *Conference) Create(name, username, startTime, endTime, password string)

```
## 通过注释生成对应接口,并实现swagger
```
func TestAutoRoute(t *testing.T) {
	err := webcmd.CreateRoute("./controllers", "项目名称", "", "")
	if err != nil {
		t.Error(err)
	}
}

```