module github.com/panyam/onehub

go 1.22

toolchain go1.22.4

require (
	github.com/jackc/pglogrepl v0.0.0-20240307033717-828fbfe908e9
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0
	github.com/panyam/dbsync v0.0.1
	github.com/panyam/goutils v0.1.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094
	gorm.io/driver/postgres v1.5.2
	gorm.io/gorm v1.25.2
)

replace github.com/panyam/dbsync v0.0.1 => ./locallinks/dbsync

// replace github.com/panyam/goutils v0.1.1 => ../goutils/
// replace github.com/panyam/slicer v0.0.1 => ../slicer/

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.4 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240624140628-dc46fd24d27d // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
