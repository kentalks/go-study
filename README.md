# A Product Manager playing with Go
### This project is about reading Go learning books.


1. Go 包中的公共方法名或公共属性，首字母必须大写，否则会被认为是私有的。Go中已经不需要public或者private
2. you can use **git merge --abort** to take the branch back to where it was in before you pulled.
3. The **git rebase** command allows you to easily change a series of commits, modifying the history of your repository.it's considered bad practice to rebase commits when you've already pushed to a repository.


## 《深入理解 Go Module》
Module可以作为Library，也可以作为调用方。
### Start a module that others can use
一个模块（module）可以包含不止一个包（package），包的用途可以单一，也可以五花八门。可以将模块[发布]()到网上，以他人应用的使用。
Go模块的下一级是包，包里再写具体的G程序。模块需要指定程序所需的依赖，比如Go的版本或者其它模块。

从模块创建开始，每一次的更新，都会发布新的模块版本。调用模块函数的开发者可以导入模块最新的包，并在更新生产环境之前用新版本进行测试。
**go mod init**命令创建一个go.mod文件，用以跟踪代码依赖。
Go按照**导出名称（exported name）**约定，只有以大写开头的函数，才可以被其它包的函数调用。
### Call your code from another module


> **Ken Xu**  
> 2022-08-23
