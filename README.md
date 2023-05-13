# 動画を起動するCLI
- netflixとDisney+をよく使うのでターミナルから起動させます

# intsall
1. cobraをインストールします
```
go install github.com/spf13/cobra-cli@latest
```
2. modファイル作成
```
go mod init go_cli
```
3. cobraの雛形作成
```
cobra-cli init
```
