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
  - リバースプロキシ
  - 静的コンテンツも配信できればベストだが一旦スコープアウト
- logging
- 設定ファイルを lua で記述する
- systemd で start できる
  - なのでコマンドとしては --foreground と --config-check のみ

### Stacks
- Go
- AWS EC2
