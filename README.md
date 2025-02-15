# cywagon

### Commands
```console
$ cywagon check -help
`check` validates config files like `nginx -t`.

$ cywagon up -help
`up` starts web server.
```

### Features
- ウェブサーバ
  - リバースプロキシ
  - 静的コンテンツも配信できればベストだが一旦スコープアウト
- logging
- 設定ファイルを lua で記述する
- lua で handler を書くことができ、リクエストパスやステータスコードを override できる

```lua
host = "example.com"

origin.host = "https://example.com"

function handler(next, req)
    if (req.path == "/favicon.ico") then
        req.path = "/aaa"
    end

    res = next(req)
    res.status = 200

    return res
end
```

### Stacks
- Go
- AWS EC2
