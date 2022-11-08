# Echo-Simple-JWT

Golang Echo で JWT を使ってみるだけ。

## 使い方

### 起動

```zsh
% make run_server
```

### リクエスト例

1. ログイン
```zsh
% curl -X POST -d 'username=admin' -d 'password=admin' http://localhost:1323/login
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY4MTU2NjM3LCJpYXQiOjE2Njc4OTc0MzcsIm5hbWUiOiJhZG1pbiJ9.TcTtep-Pdf4nvoRCUFDER0QULQqdxexRhsTQ9vun3Cs"}
```

2-1. ログイン後のリクエスト (トークンで認証が通った場合)
```zsh
% curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNjY4MTU2NzQ0LCJpYXQiOjE2Njc4OTc1NDQsIm5hbWUiOiJhZG1pbiJ9.cnQhlud42_K71Pfumr3aj8WEGhzpagL1scu_TDqQdv8" localhost:1323/restricted/hello
{"message":"Hello admin!"}
```

2-2. ログイン後のリクエスト (トークンで認証が通らなかった場合)
```zsh
% curl -H "Authorization: Bearer invalid" localhost:1323/restricted/hello
{"message":"invalid or expired jwt"}
```

## テスト

テストは関数テストと HTTP によるサーバテストがあります。

必ずサーバを起動した状態で行ってください。

### 関数テスト

```zsh
% make run_server

# 別ターミナルで
% make test
```

