# cywagon
A conditionally configurable web server. Toy app.

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

    headers = {
        "Cache-Control": "no-cache",
    }

    if {
        path = "/storage/*"

        rewrite {
            path = "/a.txt"
            // base 
        }
        respond {
            dist = "../../storage"
        }
    }

    if {
        path = "/restrict/*"
        headers_not = {"Authorization": const.basicauth}

        respond {
            status = 401
            headers = {
                "WWW-Authenticate": "Basic realm=\"Restricted\""
            }
        }
    }

    if {
        path = "/old"

        respond {
            status = 301
            headers = {
                "Location": "/",
            }
        }
    }

    if {
        path_not = "/{**/*.*,*.*}"

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
- ビルド成果物に .cywagon.lua が含まれていれば、それを尊重する

```console
$ cywagon reload <sitename>
$ cywagon check <sitename>
```
