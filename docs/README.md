# 创新实践1 项目 --- 图书管理系统

## 介绍

这是一个图书管理系统的后端， 使用 golang 完成，具体使用 gin 作为开发框架，gorm 管理数据库，数据库使用 postgre, 目前部署在本地（尚未配置docker）。默认启动在 8080 端口。

## 项目结构

项目目前分为两部分，一部分是 user 部分，负责管理用户逻辑。另一部分是 book, 负责管理图书。

整体结构图如下：

```bash
> tree            
.
├── common
│   └── user.go
├── config
├── db
│   ├── db.go
│   └── global.go
├── docs
│   └── README.md
├── go.mod
├── go.sum
├── handler
│   ├── book.go
│   ├── ping.go
│   └── user.go
├── hdu-cxsj1
├── img
│   └── book_cover
│       ├── 1.jpg
│       └── default.jpg
├── Init
│   └── init.go
├── main.go
├── middleware
│   └── userMiddle.go
├── module
│   ├── book.go
│   └── user.go
├── README.md
├── route
│   └── route.go
├── shared
│   └── consts
│       └── consts.go
└── utils
    ├── jwt
    │   └── jwt.go
    └── password
        └── hash.go

17 directories, 22 files

```

其中，db 部分负责 数据库 的更新、连接，common 负责常用的共享库，比如这里使用 GetUIDFromJWT 来从 Context 里面获取 uid。img 用于存放图片，其中， book_cover 是保存前端发来的封面图片的。utils 是一些可能用到的函数，比如密码的加密，jwt 的生成等。Init 是项目的初始化，目前为调用数据库的初始化和读取配置文件。consts 为一些常数的设计。

Module 分为 User 和 Book, 设计如下：

```go
type Book struct {
	Bid       int    `json:"bid" gorm:"primary_key" form:"bid"`
	Name      string `json:"name" gorm:"size:100;not null" form:"name"`
	Author    string `json:"author" gorm:"size:100;not null" form:"author"`
	Publisher string `json:"publisher" gorm:"size:100;not null" form:"publisher"`
	Intro     string `json:"intro" gorm:"size:255" form:"intro"`
	CoverFile string `json:"cover_file" gorm:"size:255" form:"-"`

	//User
	Uid int
}

type User struct {
	Uid      int    `json:"uid" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"unique;size:100;not null"`
	Email    string `json:"email" gorm:"unique;size:100;not null"`
	Password string `json:"password" gorm:"size:255;not null" `
	Gender   string `json:"gender" gorm:"size:10"`

	Books []Book `gorm:"foreignKey:Uid;references:Uid"`
}
```

User 外联了 Book 用于管理每个用户自己的图书。

handler 部分是主要方法的实现，同样分为 user 和 book。

Middleware 是中间件部分，主要为 JWTAuth 功能，在登陆的时候给到的 JWT token 在这里被解析。使用参数 audience 来区分权限，在这里就会阻断一部分的身份认证错误。

route 是网关部分，主要分为 ping 测试，/auth 部分用于注册登陆，/user 部分用于登陆后的用户操作，/book 用于书籍管理操作（/user /book 均需要 JWTAuth("user") 中间件管理）

