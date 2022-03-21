# CateiruSSO Go Package

CateiruSSO implementation written in go.

## Install

```bash
go get -u github.com/cateiru/cateiru-sso/pkg/go
```

## Usage

### Create Login URI

```go
CLIENT_ID := ""
RedirectURI := ""

loginUri := sso.CreateURI(CLIENT_ID, RedirectURI)
```

### Get Token and get claims

```go
TOKEN_SECRET := ""
RedirectURI := ""

token, err := sso.GetToken(code, RedirectURI, TOKEN_SECRET)
if err != nil {
    return err
}

claims, err := sso.ValidateIDToken(token.IDToken)
if err != nil {
    return err
}
```

## LICENSE

[MIT](../../LICENSE)
