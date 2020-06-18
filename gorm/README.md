# GORM

## インストール

```bash
go get github.com/jinzhu/gorm
```

MySQLのドライバが必要なのでインストール

```bash
go get github.com/go-sql-driver/mysql
```

## DBからstructureを作成する

### xoのインストール

```bash
go get -u golang.org/x/tools/cmd/goimports
go get -tags oracle -u github.com/xo/xo
```

### 生成

```bash
xo mysql://root:P@ssw0rd@127.0.0.1:13306/gormsample -o models
```
