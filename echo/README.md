# Webフレームワーク echo

## 参考サイト

https://echo.labstack.com/guide

## インストール

```bash
go get -u github.com/labstack/echo/...
```

### 確認

#### GET

```bash
curl http://localhost:1323/users 
curl http://localhost:1323/users/123
```

#### POST

```bash
curl -F "name=Hoge Smith" -F "email=hoge@mail.coom" http://localhost:1323/users
```

#### PUT

```bash
curl -F "name=Hoge Smith" -F "email=hoge@mail.coom" http://localhost:1323/users/123 -X PUT
```

## セッション

```bash
go get github.com/labstack/echo-contrib/session
```