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
- CloudFront Functions のようにリクエスト/レスポンスを整形できる
  - lua で記述

### Stacks
- Go
- AWS EC2 + Route53
- systemd とかでよしなにできたらいいなあ
- DNS サーバのセルフホストもできるけど、時間かかるので、Route 53 前提
