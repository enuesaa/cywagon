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

### features
- ウェブサーバ
- CloudFront Functions のようにリクエスト/レスポンスを整形できる
  - lua で記述

### Stacks
- Go
- EC2 でのホストを検討
- systemd とかでよしなにできたらいいなあ
- たぶん cli と worker の最低2つは必要
