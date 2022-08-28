# A Product Manager playing with Go
### This project is about reading Go learning books.
1. [Go Document](https://go.dev/doc/)
2. [Learning Go (Jon Bodner)](#)
3. [A Tour of Go](https://go.dev/tour/list)


### 零散笔记
* Go 包中的公共方法名或公共属性，首字母必须大写，否则会被认为是私有的。Go中已经不需要public或者private
* you can use **git merge --abort** to take the branch back to where it was in before you pulled.
* The **git rebase** command allows you to easily change a series of commits, modifying the history of your repository.it's considered bad practice to rebase commits when you've already pushed to a repository.


## 《深入理解 Go 模块》
模块可以作为Library，也可以作为调用方。

### `1. Start a module that others can use`
一个模块（module）可以包含不止一个包（package），包的用途可以单一，也可以五花八门。可以将模块[发布](#a-product-manager-playing-with-go)到网上，以便他人应用的使用。
模块由包组成，包里写具体的程序。模块需要指定程序所需的依赖，比如Go版本或者其它模块。

每一次模块更新（包括创建），都会发布新的模块版本。调用模块函数的开发者可以导入模块最新的包，并在更新生产环境之前用新版本进行测试。
**go mod init** 命令创建一个go.mod文件，用以跟踪代码依赖。
Go按照 **导出名称(exported name)** 约定，只有以大写开头的函数，才可以被其它包的函数调用。

### `2. Call your code from another module`
```shell
$ go mod edit -replace kensoft.tech/go-lib/greetings=../go-lib
```
若要调用本地模块，则以上的命令为带网址路径的模块指定了本地路径，然后执行**go mod tidy**以便重建新的模块依赖。  
*在VSCode中打出**helloweb**，则会快速输入go推荐的hello代码片段。*

### `3. Return and handle an error`
Handling errors is an essential feature of solid code.
**log.Fatal函**数打印错误并结束程序。

### `4. Return a random string`
通过init()函数可以在每次运行代码时改变随机值，程序启动时自动运行init函数，但在全局变量初始化后。
```go
// Init function to seed the rand package with the current time.
func init(){
    rand.Seed(time.Now().UnixNano())
}
```

### `5. Return greeting for multiple people`
**go run .**

### `6. Add a test`
测试函数的参数是testing.T类型的指针，使用该变量方法报告和记录测试情况。
*``包含的内容可理解为不转为字符串*
**go test** 命令可以运行测试函数：以Test开头的函数，且在_test.go结尾的文件中。-v标记可获得详细输出。

### `7. Compile and install the application`


> **Ken Xu**  
> 2022-08-23