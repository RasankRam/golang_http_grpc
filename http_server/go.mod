module todo-list

go 1.23.0

require (
	github.com/jackc/pgx/v5 v5.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
)

require (
	example.com/task_platform_proto v0.0.0
	github.com/go-playground/locales v0.14.1
	github.com/go-playground/universal-translator v0.18.1
	github.com/go-playground/validator/v10 v10.22.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/mdobak/go-xerrors v0.3.1
	golang.org/x/crypto v0.26.0
)

replace example.com/task_platform_proto => ../task_platform_proto

require (
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/grpc v1.67.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
