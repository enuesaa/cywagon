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
- 設定ファイルをサイト別に配置する
  - .htaccess みたいに
  - ディレクトリ単位で設定を制御できる (=反映できる) .htaccess の仕組みは優れていると感じており、vhosts のとき扱いやすいため。
  - .cywagon.lua かな
  - ディレクトリ単位はきついかな
    - どちらかというと CloudFront Functions みたいにスクリプトで制御して欲しい
  - ファイルの更新日時を見てキャッシュするイメージ
  - 共通設定も書けるように。nginx の include みたいな感じ
