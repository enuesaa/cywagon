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

## モチベーション
- 静的コンテンツを配信するWebサーバに興味がある。例えば apache や nginx など。
- ファイルを単に HTTP で serve するだけなので、仕組みとしてはシンプルだが、考えることが多い
- 往々にして設定ファイルが長くなりがち
- CloudFront Functions のような柔軟性を持ち合わせつつ宣言的に書けないか。

## TODO
- 証明書
