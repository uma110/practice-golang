# Go Practice

このリポジトリはGo言語を勉強するためのサンプルプロジェクトを格納している

主なディレクトリ:
- [echo](echo) — Echo フレームワークによるファイルアップロード例。フロントは [echo/static/index.html](echo/static/index.html) 。アップロード処理は [`file.UploadHandler`](echo/file/upload.go)（[echo/file/upload.go](echo/file/upload.go)）で実装されています。
- [todo](todo) — Gin + GORM によるシンプルな TODO アプリ（テンプレートは [todo/templates](todo/templates) 、エントリは [todo/main.go](todo/main.go)）。
- [local-server](local-server) — 静的ファイル配信と簡易 API のサンプル（[local-server/main.go](local-server/main.go)）。

実行方法（各ディレクトリで実行）:
```sh
# Echo アプリ（ファイルアップロード）
cd echo
go run main.go

# TODO アプリ
cd ../todo
go run main.go

# ローカル静的サーバー
cd ../local-server
go run main.go
```

補足:
- Echo のアップロードはフォーム（[echo/static/index.html](echo/static/index.html)）から `/upload` に POST され、処理は [`file.UploadHandler`](echo/file/upload.go) が担当します。
- Echo サンプルは Supabase Storage を利用するため、必要に応じて `echo/file/upload.go` 内の `projectId` / `apiKey` 等を調整してください。
- 各モジュールで依存解決が必要な場合は、該当ディレクトリで `go mod tidy` を実行してください。

参考ファイル:
- [echo/main.go](echo/main.go)
- [echo/file/upload.go](echo/file/upload.go)
- [echo/static/index.html](echo/static/index.html)
- [todo/main.go](todo/main.go)
- [todo/templates/index.html](todo/templates/index.html)
- [local-server/main.go](local-server/main.go)