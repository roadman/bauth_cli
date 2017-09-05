# bauth_cli

Basic Auth Setting Generate Command.

## Description

## Usage

bauth_cli generate login num

```bash
$ bauth_cli generate bauth 10
# group: bauth, created: 2017-09-05 14:55:41.183805689 +0900 JST
# login: bauth001 , pass: s5nhgyrv
bauth001:$apr1$s/b2wrto$Ja8FceE/js6cC8g3UKkSS0
# login: bauth002 , pass: glmk9byh
bauth002:$apr1$eMIJaVBY$ZtRA9pp9FK1mdaYaFWssm/
# login: bauth003 , pass: hu9sst4q
bauth003:$apr1$XiCJcMcw$psgJe6ptxOksRu4v73uY0/
# login: bauth004 , pass: edw89r4d
bauth004:$apr1$AKIHbU12$Dj039lJQoMqy3zAXm9Po5.
# login: bauth005 , pass: 5mw2t1n3
bauth005:$apr1$69HELFgN$xsLl6JuG9ahx5DsVJlliu/
# login: bauth006 , pass: 7pncsmnj
bauth006:$apr1$QBnnMK31$oMf5WU4pVi6X8I8wyrbNC.
# login: bauth007 , pass: jtvzpqcz
bauth007:$apr1$19wsVrnm$0fHBBnvA15iiacN0KBI991
# login: bauth008 , pass: f3g91x66
bauth008:$apr1$DCzNL3tG$Qap5K8rfVWuq5SsTzhsT6/
# login: bauth009 , pass: 9ll9mp3f
bauth009:$apr1$qCjH0JtD$bRogeaWSRRYIC0hLALJYs0
# login: bauth010 , pass: p0ihuega
bauth010:$apr1$RovOAFck$6/.kPYDRqnhBSoNZkavJq/
$
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/roadman/bauth_cli
$ cd $GOPATH/src/github.com/roadman/bauth_cli
$ go build
```

## Contribution

1. Fork ([https://github.com/roadman/bauth_cli/fork](https://github.com/roadman/bauth_cli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[roadman](https://github.com/roadman)
