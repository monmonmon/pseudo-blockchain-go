# pseudo-blockchain-go
Golang と gRPC の演習問題として、こちらの記事を参考に実装。
[gRPCとgolangでブロックチェーンを作ってみる - こさちゃん技術アップブログ](http://suga-tech3.hatenablog.com/entry/2018/01/22/210616)

gRPC のパッケージをインストール
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

.proto ファイルから .go ファイルを生成
```
protoc --go_out=plugins=grpc:. proto/blockchain.proto
```

サーバ起動。8080 で待ち受ける
```
go run server/main.go
```

クライアントからブロック追加をリクエスト
```
go run client/main.go --add
```

クライアントからブロック一覧の取得をリクエスト
```
go run client/main.go --list
```
