# 動画を起動するCLI
- netflixとDisney+をよく使うのでターミナルから起動させます
# 起動コマンド
- netflixの起動コマンド
```
video n
```
- disney+の起動コマンド
```
video d
```
  - フラグ
    - -mをつけることでマーベル映画を限定
# intsall
1. cobraをインストールします
```
go install github.com/spf13/cobra-cli@latest
```
2. modファイル作成
```
go mod init video
```
3. cobraの雛形作成
```
cobra-cli init
```
4. 以下のコマンドを打つ
```
cobra-cli add n && cobra-cli add d
```
5. cmd以下のディレクトリをコピー
