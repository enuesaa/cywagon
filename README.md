# cywagon

## Planning
### Commands
```console
$ cywagon create
failed to create project due to admin server does not started.
please run `cywagon up`

$ cywagon up

$ cywagon create <project-name> --hostname <hostname>

$ cywagon ls --filter <project-name>:<prefix>
<project-name>:<version-name> published
<project-name>:<version-name>

$ cywagon push <project-name>:<version-name> --from-dir .
<project-name>:<version-name>

$ cywagon publish <project-name>:<version-name>

$ cywagon down
```

### Features
- ウェブサーバ
  - 静的コンテンツのウェブサーバ
  - 将来的には PHP など動かしたいが実装や FastCGI への理解に時間がかかりそうなので、一旦は静的コンテンツのみ
- 設定ファイルを lua で記述する
- handler も lua で記述できる
  - CloudFront Functions のようにリクエスト/レスポンスを整形できる
  - CloudFront はリクエスト/レスポンスとトリガーが分かれているが、挙動が案外わかりづらいので、いわゆるミドルウェアみたいに next() を呼ぶ方式にする。
- systemd で start できる
  - なのでコマンドとしては --foreground と --config-check のみ

### Stacks
- Go
- AWS EC2
