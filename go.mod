module github.com/danicc097/openapi-go-gin-postgres-sqlc

go 1.20

require (
	github.com/danicc097/oidc-server/v3 v3.1.0
	github.com/danicc097/xo/v5 v5.1.0
	github.com/dave/dst v0.27.3
	github.com/deepmap/oapi-codegen v1.16.2
	github.com/deepmap/oapi-codegen/v2 v2.0.1-0.20231204155340-1f53862bcc64
	github.com/fatih/structs v1.1.0
	github.com/fatih/structtag v1.2.0
	github.com/getkin/kin-openapi v0.122.0
	github.com/gin-contrib/cors v1.5.0
	github.com/gin-contrib/pprof v1.4.0
	github.com/gin-contrib/zap v0.2.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-jet/jet/v2 v2.10.1
	github.com/go-playground/validator/v10 v10.15.5
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang-migrate/migrate/v4 v4.16.2
	github.com/google/go-cmp v0.5.9
	github.com/google/uuid v1.4.0
	github.com/iancoleman/strcase v0.3.0
	github.com/jackc/pgconn v1.14.1
	github.com/jackc/pgerrcode v0.0.0-20220416144525-469b46aa5efa
	github.com/jackc/pgx-zap v0.0.0-20221202020421-94b1cb2f889f
	github.com/jackc/pgx/v5 v5.5.0
	github.com/joho/godotenv v1.5.1
	github.com/kenshaw/inflector v0.2.0
	github.com/kenshaw/snaker v0.2.0
	github.com/oapi-codegen/runtime v1.1.0
	github.com/oapi-codegen/testutil v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.17.0
	github.com/smartystreets/assertions v1.13.1
	github.com/smartystreets/gunit v1.4.5
	github.com/stretchr/testify v1.8.4
	github.com/swaggest/jsonschema-go v0.3.62
	github.com/swaggest/openapi-go v0.2.42
	github.com/zitadel/oidc/v2 v2.11.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.45.0
	go.opentelemetry.io/contrib/propagators/jaeger v1.20.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	go.opentelemetry.io/otel/trace v1.19.0
	go.uber.org/zap v1.26.0
	golang.org/x/exp v0.0.0-20230728194245-b0cb94b80691
	golang.org/x/sync v0.4.0
	golang.org/x/text v0.14.0
	mvdan.cc/gofumpt v0.5.0
)

require (
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/Joker/jade v1.1.3 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/alicebob/gopher-json v0.0.0-20230218143504-906a9b012302 // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.10.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gomarkdown/markdown v0.0.0-20230716120725-531d2d74bc12 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/iris-contrib/schema v0.0.6 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.2 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/kataras/blocks v0.0.7 // indirect
	github.com/kataras/golog v0.1.9 // indirect
	github.com/kataras/iris/v12 v12.2.6-0.20230908161203-24ba4e8933b9 // indirect
	github.com/kataras/pio v0.0.12 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/mailgun/raymond/v2 v2.0.48 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/microcosm-cc/bluemonday v1.0.25 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/muhlemmer/gu v0.3.1 // indirect
	github.com/muhlemmer/httpforwarded v0.1.0 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/prometheus/client_model v0.4.1-0.20230718164431-9a2bf3000d16 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.11.1 // indirect
	github.com/rs/cors v1.10.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/swaggest/refl v1.3.0 // indirect
	github.com/tdewolff/minify/v2 v2.12.9 // indirect
	github.com/tdewolff/parse/v2 v2.6.8 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	github.com/yuin/gopher-lua v1.1.0 // indirect
	github.com/zitadel/logging v0.3.4 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
)

require (
	github.com/alicebob/miniredis/v2 v2.30.5
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gin-contrib/sse v0.1.0
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/labstack/echo/v4 v4.11.3 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/lib/pq v1.10.9
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.43.0
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/time v0.3.0
	golang.org/x/tools v0.14.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
)
