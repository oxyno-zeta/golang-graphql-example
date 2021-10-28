module github.com/oxyno-zeta/golang-graphql-example

go 1.13

require (
	cirello.io/pglock v1.8.0
	github.com/99designs/gqlgen v0.14.0
	github.com/99designs/gqlgen-contrib v0.1.1-0.20200601100547-7a955d321bbd
	github.com/AppsFlyer/go-sundheit v0.5.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/gzip v0.0.3
	github.com/gin-contrib/static v0.0.1
	github.com/gin-gonic/gin v1.7.4
	github.com/go-gormigrate/gormigrate/v2 v2.0.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/golang/mock v1.6.0
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/opentracing-contrib/go-gin v0.0.0-20201220185307-1dd2273433a4
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/pquerna/cachecontrol v0.1.0 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	github.com/thoas/go-funk v0.9.1
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/vektah/gqlparser/v2 v2.2.0
	github.com/xhit/go-simple-mail/v2 v2.10.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gorm.io/driver/postgres v1.1.1
	gorm.io/driver/sqlite v1.1.1
	gorm.io/gorm v1.21.15
	gorm.io/plugin/opentracing v0.0.0-20210706093620-707e98269c0e
	gorm.io/plugin/prometheus v0.0.0-20210820101226-2a49866f83ee
)

replace github.com/99designs/gqlgen-contrib => github.com/oxyno-zeta/gqlgen-contrib v0.1.1-0.20210822164044-9b33d2c27fa1
