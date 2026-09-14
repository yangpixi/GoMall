module github.com/yangpixi/GoMall/gateway

go 1.26.5

require github.com/yangpixi/GoMall/shared v0.0.0

require (
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
)

replace github.com/yangpixi/GoMall/shared v0.0.0 => ../shared
