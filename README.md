

# linweb

> #### linweb是一套简单自由的web框架，适合一些如博客一样的简单web系统，不过度追求性能，注重开发的便捷、整洁及代码的扩展性。



## 面向接口编程

在linweb中，将完全面向接口编程并将可扩展部分插件化。

linweb提供一套插件接口及默认实现，你也可以通过***AddCustomizePlugins***方法添加自定义实现。

<img src=".\docs\images\structure.png" alt="image-20210727102845643" style="zoom:80%;" />



## 接口即文档

插件接口都**在/interfaces文件目录下**并有尽量详细的注释，可以通过对接口方法的查看理解并应用linweb。



## 如何使用linweb？

> ###### 详细示例都在examples目录下
>

### Run

使用NewLinweb方法创建一个linweb，调用Run方法就可以运行一个没有任何api的web项目。

```go
func main() {
	l := linweb.NewLinweb()
	l.Run(":9999")
}
```

### Controller

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