<!-- omit in toc -->
# treezoo_frontend

動物園で暮らす動物たちの家系図アプリ「TreeZoo」のバックエンドを提供するGoプロジェクトです。フロントエンドを提供するFlutterプロジェクトは [treezoo_frontend](https://github.com/yutaiwamoto1114/treezoo_frontend) にて公開しています。


- [1. 構成](#1-構成)
- [2. ディレクトリ構成](#2-ディレクトリ構成)
- [3. デバッグ方法](#3-デバッグ方法)
- [4. ビルド方法](#4-ビルド方法)
- [5. デプロイ](#5-デプロイ)
- [6. Goバージョンアップ](#6-goバージョンアップ)

## 1. 構成
- Go 1.22.3

## 2. ディレクトリ構成
プロジェクトの規模が拡大するにつれて、[golang-standards](https://github.com/golang-standards/project-layout/blob/master/README_ja.md#standard-go-project-layout) を参考にディレクトリを構成していきます。
```
.
├── main.go: アプリケーションのエントリーポイントであり、APIルーティングを定義します。
├── api: APIエンドポイントを定義します。
├── docs: Swaggerドキュメントを管理します。
├── db: データベース接続とその設定を管理します。
├── model: データベース操作に利用するモデルオブジェクトを管理します。
├── sql: データベースのテーブル・ビュー作成に利用したSQLクエリを保管します。
└── under_consideration: 検討中やテスト中のソースです。将来的に削除します。
```

## 3. デバッグ方法
`go run .\main.go`

## 4. ビルド方法
`go build -o treezoo-backend.exe main.go`

環境切替方法は未検討です。
- dev環境
- stg環境
- pord環境

## 5. デプロイ
未検討です。

## 6. Goバージョンアップ
1. 現在のバージョンを確認します。
    `go version`
2. [All releases - The Go Programming Language](https://go.dev/dl/) にて、最新のStableバージョンを確認します。
3. 最新バージョンのGoをダウンロードします。
    `go1.22.3 download`
4. プロジェクト直下の`go.mod`に記載されているGoバージョンを以下のコマンドで修正します。
    `go mod edit -go="1.22.3"`