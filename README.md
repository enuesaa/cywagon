# cywagon
A conditionally configurable web server. Toy app

[![ci](https://github.com/enuesaa/cywagon/actions/workflows/ci.yml/badge.svg)](https://github.com/enuesaa/cywagon/actions/workflows/ci.yml)

### Commands
```console
$ cywagon check -help
`check` validates config files like `nginx -t`.

$ cywagon up -help
`up` starts web server.
```

### Features
- 静的コンテンツを配信するWebサーバ
- 設定ファイルを HCL で記述する

```hcl
site "sampleapp" {
  host = "localhost:3000"
  dist = "../dist"

  if {
    path = "/old"

    respond {
      status = 301
      headers = {
        "Location" : "/",
      }
    }
  }

  if {
    path_not_in = ["/**/*.*", "/*.*"]

    rewrite {
      path = "/index.html"
    }
  }
}
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
