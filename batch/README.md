# クローリングアプリ

指定のURLをクローリングし、DBに記録する。

- 別ドメインのリンクは保存しない。

## setup

```
//.ent
go get -d entgo.io/ent/cmd/ent

//MySQLドライバ
go get github.com/go-sql-driver/mysql

//スキーマの初期化
go run entgo.io/ent/cmd/ent init CrawledUrl LinkUrl

//もろもろファイル生成
go generate ./ent

//マイグレーション実行確認
docker-compose exec db mysql -uroot -p demo

//dotenv
go get github.com/joho/godotenv
```