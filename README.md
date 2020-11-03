├── go.mod
├── go.sum
├── gqlgen.yml               - gqlgen構成ファイル、生成されたコードを制御するためのノブ。
├── graph
│   ├── generated            - 生成されたランタイムのみを含むパッケージ
│   │   └── generated.go
│   ├── model                - 生成された、またはその他のすべてのグラフモデルのパッケージ
│   │   └── models_gen.go
│   ├── resolver.go          - ルートグラフリゾルバタイプ。このファイルは再生成されません
│   ├── schema.graphqls      - いくつかのスキーマ。スキーマを必要な数のgraphqlファイルに分割できます
│   └── schema.resolvers.go  - schema.graphqlのリゾルバー実装
└── server.go                - アプリへのエントリポイント。必要に応じてカスタマイズしてください
