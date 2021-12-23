# About This Template

- [x] gin, [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [x] systray, [https://github.com/getlantern/systray](https://github.com/getlantern/systray)
- [X] air, [https://github.com/cosmtrek/air](https://github.com/cosmtrek/air)
- [x] gorm, [https://gorm.io/](https://gorm.io/)
- [x] swagger, [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- [x] jwt, [https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [x] ini, [https://github.com/go-ini/ini](https://github.com/go-ini/ini)

## Environment

- Windows 10 x64
- Go 1.17
- NPM v8 (from Node.js)
- Git Bash

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

- `ready:swag`: generate swagger files to folder `/swagger`
- `serve:air`: run air for development
- `serve:certs`: generate self-signed certificate to folder `/air` (development)
- `build:certs`: generate self-signed certificate to folder `/build` (production)
- `build:run`: compile App to folder `/build` with filename `App.exe`, then run it
- `build`: compile App to folder `/build` with filename `App.exe` and ldflags `-H=windowsgui`

## Path Structure

- `/air`: generated files for development
  - `.air.toml`: air cli config
- `/build`: generated files for production
- `/logger`: logger for `models`, `routers` and `tray`
- `/models`: database connection and table models
- `/routers`: define server and route groups
- `/swagger`: generated files for swagger
- `/tray`: define system tray

> `main.config.go` - all configs combined for the whole App

> Other `*.config.go` files are configs for different packages

