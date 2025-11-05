lyrics-cli

Go言語とCobraライブラリで作成された、歌詞検索CLIツールです。
Lyrics.ovh API を使用して、指定されたアーティストと曲名の歌詞をターミナルに表示します。

概要

このプロジェクトは、Go言語の学習（net/http, json, cobra）の一環として作成されました。

cobra を使ったCLIの骨組み

net/http を使った外部APIへのリクエスト

encoding/json を使ったレスポンスのパース

net/url を使ったURLエンコード

などの基本的な要素を含んでいます。

インストール

go install を使ってインストールできます。（$GOPATH/bin または $HOME/go/bin にPATHが通っている必要があります）

go install [github.com/](https://github.com/)shunshun0803/lyrics-cli@latest


または、リポジトリをクローンしてローカルでビルドすることも可能です。

git clone [https://github.com/](https://github.com/)shunshun0803/lyrics-cli.git
cd lyrics-cli
go build -o lyrics-cli .
# ./lyrics-cli search ... で実行


使い方

search サブコマンドを使用します。アーティスト名と曲名を引数として渡してください。

書式:

lyrics-cli search "[アーティスト名]" "[曲名]"


実行例:

$ lyrics-cli search "Queen" "Bohemian Rhapsody"

Is this the real life?
Is this just fantasy?
Caught in a landside,
No escape from reality
...
(以下、歌詞が続く)


歌詞が見つからない場合

$ lyrics-cli search "Artist" "Unknown Song"

Error: 歌詞が見つかりませんでした。
APIからのエラー: No lyrics found
