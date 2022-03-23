

# linweb

> #### linweb是一套简单自由的web框架，适合一些如博客一样的简单web系统，不过度追求性能，注重开发的便捷、整洁及代码的扩展性。



## 接口即文档

插件接口都**在/interfaces文件目录下**并有尽量详细的注释，你只需要查看interfaces目录下的文件接口就可以了解到linweb的功能。

linweb目前实现了请求上下文（context）、动态路由（router）、中间件（middleware）、模型验证与模型映射（model --- validate、map）等，其中部分实现参考了 [极客兔兔的七天实现Web框架](https://github.com/geektutu/7days-golang) 。待开发功能及工作可在[Roadmap](https://github.com/Codexiaoyi/linweb/issues/1)中查看，欢迎建议、issue、pr和star~

###### PS：目前该项目仅作为玩具，接受大佬建议并完善，不喜勿喷。



## 面向接口编程

在linweb中，将完全面向接口编程并将可扩展部分插件化。

linweb提供一套插件接口及默认实现，你也可以通过***AddCustomizePlugins***方法添加自定义实现。

<img src=".\docs\images\structure.png" alt="image-20210727102845643" style="zoom:80%;" />



## 如何使用linweb？

> ###### 详细示例都在examples目录下

- ### Run


使用NewLinweb方法创建一个linweb，调用Run方法就可以运行一个没有任何api的web项目。

```go
func main() {
	l := linweb.NewLinweb()
	l.Run(":9999")
}
```

- ### Controller


linweb将面向Controller定义api接口。

#### 1.你需要在根目录下建立/controllers目录（待优化，目前是必须需要在controllers目录下）

<img src=".\docs\images\controllers.png" alt="image-20210727111727506" style="zoom:150%;" />

#### 2.定义controller

①需要在controller方法的注释中**添加注解，标识HTTP方法和路由路径**。如果没有，将不作为一个http请求接口。

②**方法的第一个参数必须为IContext**，linweb将自动实例化，Context中保存request及response的信息。

③如果存在dto入参，linweb将自动解析request.body的json字符串，并将其转化为dto实例。

```go
type LoginDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type UserController struct {
}

//[GET("/hello")]
func (user *UserController) Hello(c interfaces.IContext) {
	c.Response().HTML(http.StatusOK, "<h1>Hello linweb</h1>")
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	fmt.Println(dto)
	c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
}

```

#### 3.注册所有的controller到linweb中

```go
func main() {
	l := linweb.NewLinweb()
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	l.Run(":9999")
}
```

- ### Cache

框架内置了缓存，可以自定义最大内存占用，达到上限采用**LRU-K淘汰策略**。缓存可以设置键有效时间，并使用 **惰性删除 和 定期删除** 策略，实现对过期键及时扫描删除。我们可以直接使用linweb提供的Cache单例。

```go
type BlogController struct {
}

//[GET("/blog/:id")]
func (blog *BlogController) GetBlog(c interfaces.IContext) {
	id, _ := strconv.Atoi(c.Request().Param("id"))
	linweb.Cache.AddWithExpire("id", id, 10*time.Second)
	c.Response().String(http.StatusOK, "id=%s", c.Request().Param("id"))
}
```

- ### Middleware

使用***AddMiddlewares***方法添加多个针对所有api接口的全局中间件。

```go
func main() {
	l := linweb.NewLinweb()
	l.AddMiddlewares(PrintHelloMiddleware)
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	l.Run(":9999")
}

func PrintHelloMiddleware(c interfaces.IContext) {
	fmt.Println("hello linweb!")
	c.Next()
	fmt.Println("byebye linweb")
}
```

- ### Plugins

可以通过***AddCustomizePlugins***方法添加自定义的插件（插件必须实现相应的接口），未添加的插件将使用默认插件。

```go
func main() {
	l := linweb.NewLinweb()
	l.AddCustomizePlugins(&CustomizeModel{},&CustomizeRouter{})
	l.AddControllers(&controllers.UserController{}, &controllers.BlogController{})
	l.Run(":9999")
}
```

- ### Model

model用于对struct的模型验证与映射，采用**链式调用**的方式，可以通过***NewModel***方法传入需要操作的struct，通过调用Validate、MapToByFieldName等方法实现验证和Dto映射。目前使用[validator](https://github.com/go-playground/validator)、[go-mapper](https://github.com/Codexiaoyi/go-mapper)实现验证与映射功能，使用规则详见链接。

```go
type LoginDto struct {
	Name     string `json:"Name"`
	Password string `json:"Password"`
}

type DatabaseModel struct {
	Name     string
	Password string
}

type UserController struct {
}

//[POST("/login")]
func (user *UserController) Login(c interfaces.IContext, dto LoginDto) {
	dataModel := &DatabaseModel{}
	err := linweb.NewModel(dto).Validate().MapToByFieldName(dataModel).ModelError()
	if err != nil {
		c.Response().String(http.StatusInternalServerError, "Model error :%s!", err.Error())
	}
	c.Response().String(http.StatusOK, "Welcome %s!", dto.Name)
}
```

