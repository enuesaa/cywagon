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

### メモ
containerd の systemd unit file
- https://raw.githubusercontent.com/containerd/containerd/main/containerd.service

### Future plan
- コンテナベースにしない
- 静的コンテンツを配信するのみ
  - もともとこちらに興味があったため
- インメモリ or 内部的に静的コンテンツを格納する
- CloudFront + S3 みたいにパブリッシュ & キャッシュ削除できるイメージ
- ビルド成果物に .cywagon.lua が含まれていれば、それを尊重する
- 結局 SSG だと dist ディレクトリが必要になるので、apache 風味にサイト単位のディレクトリは作らない
- いったん認証認可なし

```console
$ cywagon deploy <sitename> .
{
  deploymentId: <id>
}
$ cywagon publish <sitename>:<deploymentId>
```
