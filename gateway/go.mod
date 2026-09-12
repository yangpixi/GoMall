module github.com/yangpixi/GoMall/gateway

go 1.26.5

require (
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/yangpixi/GoMall/shared v0.0.0
)

replace github.com/yangpixi/GoMall/shared v0.0.0 => ../shared
