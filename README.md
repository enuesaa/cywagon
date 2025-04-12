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

### Future plan
- コンテナベースにしない
- 静的コンテンツを配信するのみ
  - もともとこちらに興味があったため
- インメモリに静的コンテンツを格納する
- CloudFront + S3 みたいにパブリッシュ & キャッシュ削除できるイメージ
- 理想としては HTTP でデプロイできるようにしたいが、認証認可を整備する余裕がない
  - そのため特定のディレクトリへ静的コンテンツを置き reload する方式へ
  - ただし apache みたいに即座に反映されるのではなく、明示的に reload する
    - インメモリで格納しているため。
    - こちらの方が「配置~デプロイ」まで時間的猶予が生まれ、安全であるため。
- ビルド成果物に .cywagon.lua が含まれていれば、それを尊重する

```console
$ cywagon reload <sitename>
$ cywagon check <sitename>
```
