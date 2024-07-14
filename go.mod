module github.com/panyam/onehub

go 1.22

toolchain go1.22.4

require (
	github.com/jackc/pglogrepl v0.0.0-20230826184802-9ed16cb201f6
	github.com/lib/pq v1.10.9
	github.com/typesense/typesense-go v0.8.0 // indirect
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.1
	github.com/panyam/dbsync v0.0.1
	github.com/panyam/goutils v0.0.40
	github.com/stretchr/testify v1.8.4
	google.golang.org/genproto/googleapis/api v0.0.0-20230724170836-66ad5b6ff146
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.2
)

replace github.com/panyam/dbsync v0.0.1 => ./locallinks/dbsync

// replace github.com/panyam/goutils v0.0.37 => ../goutils/
// replace github.com/panyam/slicer v0.0.1 => ../slicer/

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/deepmap/oapi-codegen v1.12.3 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sony/gobreaker v0.5.0 // indirect
	golang.org/x/crypto v0.11.0 // indirect
	golang.org/x/net v0.12.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/text v0.11.0 // indirect
	google.golang.org/genproto v0.0.0-20230706204954-ccb25ca9f130 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230724170836-66ad5b6ff146 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
