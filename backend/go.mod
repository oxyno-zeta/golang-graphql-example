module github.com/oxyno-zeta/golang-graphql-example

go 1.21

toolchain go1.21.3

require (
	cirello.io/pglock v1.14.1
	emperror.dev/errors v0.8.1
	github.com/99designs/gqlgen v0.17.39
	github.com/99designs/gqlgen-contrib v0.1.1-0.20230605152740-d49a0a0239a7
	github.com/AppsFlyer/go-sundheit v0.5.0
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/coreos/go-oidc/v3 v3.7.0
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/fsnotify/fsnotify v1.6.0
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-contrib/gzip v0.0.6
	github.com/gin-contrib/static v0.0.1
	github.com/gin-gonic/gin v1.9.1
	github.com/go-gormigrate/gormigrate/v2 v2.1.1
	github.com/go-playground/validator/v10 v10.15.5
	github.com/gofrs/uuid v4.4.0+incompatible
	github.com/graph-gophers/dataloader/v7 v7.1.0
	github.com/hasura/go-graphql-client v0.10.0
	github.com/prometheus/client_golang v1.17.0
	github.com/rabbitmq/amqp091-go v1.9.0
	github.com/ravilushqa/otelgqlgen v0.13.1
	github.com/spf13/viper v1.17.0
	github.com/stretchr/testify v1.8.4
	github.com/thoas/go-funk v0.9.3
	github.com/vektah/gqlparser/v2 v2.5.10
	github.com/xhit/go-simple-mail/v2 v2.16.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.45.0
	go.opentelemetry.io/contrib/propagators/b3 v1.20.0
	go.opentelemetry.io/contrib/propagators/jaeger v1.20.0
	go.opentelemetry.io/contrib/propagators/ot v1.20.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.19.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/trace v1.19.0
	go.uber.org/automaxprocs v1.5.3
	go.uber.org/goleak v1.2.1
	go.uber.org/mock v0.3.0
	go.uber.org/zap v1.26.0
	golang.org/x/oauth2 v0.13.0
	gorm.io/driver/postgres v1.5.3
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gorm v1.25.5
	gorm.io/plugin/opentelemetry v0.1.4
	gorm.io/plugin/prometheus v0.0.0-20231010074352-de33ebe52539
)

require (
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.10.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-test/deep v1.1.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.17 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/sagikazarmark/locafero v0.3.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sosodev/duration v1.2.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/toorop/go-dkim v0.0.0-20201103131630-e1cd1a0a5208 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	go.opentelemetry.io/contrib v1.20.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)
