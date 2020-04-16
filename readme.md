# go-atcoder-workspace

atcoderにGoで参加するための設定がセットアップ済みのリポジトリです。  
現時点ではLinuxとMacに対応しています。  
Windows向けには、WSLとDockerでのセットアップ手順を今後追加予定です。

## 特徴
* `make new`によるコンテストごとのソースコードテンプレート自動生成(powered by [atcoder-tools](https://github.com/kyuridenamida/atcoder-tools))
* `make submit`によるソースコードの提出(powered by [atcoder-tools](https://github.com/kyuridenamida/atcoder-tools))
* 複数ファイル実装のサポート
* ジェネリクスとコード自動生成による効率的なライブラリ実装環境の提供

## 準備
以下をあらかじめインストールしておいてください。
* Go (実装したコードをコンパイルするのに必要です。[genny](https://github.com/cheekybits/genny), [mustify](https://github.com/mpppk/mustify), [gollup](https://github.com/mpppk/gollup)のインストールにも利用します。)
* python3.5以上 ([atcoder-tools](https://github.com/kyuridenamida/atcoder-tools)を動かすのに必要です)
* pip3

## Setup
```shell
$ git clone https://github.com/mpppk/go-atcoder-workspace
$ cd go-atcoder-workspace
$ make setup # atcoder-tools, genny, mustify, gollupをインストールします
$ make generate # ライブラリコードの生成を行います
```
直接git cloneするのでなく、Forkしてから使うとそのままコンテスト用のコードをコミットできて便利かもしれません。

### Note
2020/4/16時点ではatcoder-toolsにGoサポートのパッチがマージされていないので、代わりにforkされたリポジトリからインストールします。
```shell
$ git clone https://github.com/nu50218/atcoder-tools
$ cd atcoder-tools
$ pip3 install -e .
```

### コンテスト実施の流れ
abc999に参加する設定で、コンテストへの参加からサブミットまでを説明します。

1. コンテストを開始するには`go-atcoder-workspace`のルートディレクトリで`make new contest=abc999`を実行します
    * `abc999`というディレクトリが生成されます
    * `abc999`以下にA~Fまでの各設問に対応するディレクトリが生成されます
    * 各設問ディレクトリの`main.go`には、設問内容に応じて入力を受け取るコードがあらかじめ実装されています。
        * 例として`make new contest=abc126`を実行して生成されたディレクトリをリポジトリにコミットしてあります。
1. ライブラリの変更が必要な場合は、`lib`以下のファイルを変更してから`make generate`を実行します。
1. `in_*.txt`と対応する`out_*.txt`を追加することで、テストを追加できます。
1. `make test pkg=abc999/A`を実行すると、提出用のコードを生成し、設問の入力例に応じたテストが実行されます。
1. `make submit pkg=abc999/A`を実行すると、提出用のコードを生成し、提出します。
    * submit前にtestが自動で実行されます。失敗した場合は提出が中断されるので安心です。
    * 例として`abc126/A/submiit/submit.go`に提出用に生成されたコードをコミットしてあります。

## ディレクトリ構成

* `lib` 汎用的に利用するライブラリのコードをおきます。コードにgennyを利用した場合、そのままでは使えません。`make generate`を実行して実際に利用するコードを生成してください。
* `templates` `make new`で生成する`main.go`のテンプレートを置きます。

## 複数ファイルによる実装
atcoderでは複数ファイルから成るソースコードを提出することはできません。
go-atcoder-workspaceでは、実装に利用したファイルをマージした提出用ソースコードを生成することで、この問題を回避できます。
各設問の`main.go`からは、同じmain packageかlibディレクトリ以下のlib packageを利用できます。  
例えば、以下のコードがあるとします。

```go
func solve(N []int) int {
    return lib.MaxInt(N...)
}
```

`make test`や`make submit`を実行すると、提出用のコードとして以下のような`submit/submit.go`が生成されます。

```go
func lib_MaxInt() {
    ...
}

func solve(N []int) int {
    return lib.MaxInt(N...)
}
```

## ライブラリ実装について
記述量を減らし、効率的にライブラリを実装できるようにgennyとmustifyという2つのコード自動生成ツールを利用しています。
* (TODO) gennyについて
* (TODO) mustifyについて

## Docker
* (TODO) `docker run -it mpppk/go-atcoder-workspace`で開発が始められるDocker imageを作って説明を書く
* (TODO) docker-compose

## Visual Studio Codeでの例
* (TODO) VSCode Remove Development+Dockerでの開発手順
* (TODO) VSCode Remove Development+WSLでの開発手順
