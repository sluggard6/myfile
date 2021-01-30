module github.com/sluggard/myfile

//module myfile

go 1.15

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/fatih/structs v1.1.0
	github.com/go-delve/delve v1.5.1 // indirect
	github.com/go-openapi/spec v0.20.1 // indirect
	github.com/go-openapi/swag v0.19.13 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/google/uuid v1.1.5
	github.com/iris-contrib/swagger/v12 v12.0.1
	github.com/kataras/iris/v12 v12.2.0-alpha2.0.20210110101619-f4989bd5aaac
	github.com/sirupsen/logrus v1.7.0
	github.com/swaggo/swag v1.7.0
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.0.0-20210115202250-e0d201561e39 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.11
// gorm.io/driver/sqlmock v0.0.0
)

// replace gorm.io/driver/sqlmock => ../sqlmock
