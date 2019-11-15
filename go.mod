module srencat

go 1.13

require (
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/thedevsaddam/renderer v1.2.0 // indirect
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	srencat/mvc_controller v0.0.1
	srencat/mvc_model v0.0.1
	srencat/mvc_route v0.0.1
)

replace srencat/mvc_controller v0.0.1 => ./controller

replace srencat/mvc_route v0.0.1 => ./router

replace srencat/mvc_model v0.0.1 => ./model
