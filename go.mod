module github.com/fuloge/basework

go 1.14

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/elastic/go-elasticsearch/v7 v7.7.0
	github.com/gin-gonic/gin v1.6.2
	github.com/gomodule/redigo v1.8.1
	github.com/kr/pretty v0.2.0 // indirect
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	go.uber.org/zap v1.14.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	xorm.io/xorm v1.0.3 // indirect

)

replace (
	xorm.io/xorm v1.0.3 => gitea.com/xorm/xorm v1.0.3
)
