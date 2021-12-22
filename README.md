# About This Template

- [x] gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [x] systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- [ ] air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- [ ] gorm, [https://gorm.io/](https://gorm.io/)
- [x] swagger, [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- [ ] jwt, [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)

## Environment

- Windows 10 x64
- Go 1.17
- NPM v8 (from Node.js)

## Setup

``` bash
# install air cli
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
# install swag cli
go install github.com/swaggo/swag/cmd/swag@latest
# install go dependencies
go mod tidy
```

## Scripts

- `build:certs`: create self-signed certificate for TLS
- `build:swag`: generate swagger files to folder `/swagger`
- `build:run`: compile App to folder `/build` with filename `App.exe`

