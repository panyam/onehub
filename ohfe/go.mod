module github.com/panyam/onehub/ohfe

go 1.22

toolchain go1.22.4

// replace github.com/panyam/s3gen v0.0.11 => ../../s3gen/

// replace github.com/panyam/goutils v0.1.1 => ../../goutils/

// replace github.com/panyam/slicer v0.0.1 => ../slicer/

require (
	github.com/alexedwards/scs/v2 v2.8.0
	github.com/felixge/httpsnoop v1.0.4
	github.com/gorilla/mux v1.8.1
	github.com/panyam/goutils v0.1.2
	github.com/panyam/onehub v0.0.0-20240625081306-c666a276c24b
	github.com/panyam/s3gen v0.0.11
	golang.org/x/text v0.14.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/adrg/frontmatter v0.2.0 // indirect
	github.com/alecthomas/chroma/v2 v2.14.0 // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/morrisxyang/xreflect v0.0.0-20231001053442-6df0df9858ba // indirect
	github.com/radovskyb/watcher v1.0.7 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/yuin/goldmark v1.7.1 // indirect
	github.com/yuin/goldmark-highlighting/v2 v2.0.0-20230729083705-37449abec8cc // indirect
	go.abhg.dev/goldmark/anchor v0.1.1 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/gorm v1.25.2 // indirect
)
