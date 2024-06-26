module github.com/panyam/onehub/ohfe

go 1.22

toolchain go1.22.4

replace github.com/panyam/s3gen v0.0.6 => ../../s3gen/

// replace github.com/panyam/slicer v0.0.1 => ../slicer/

require (
	github.com/alexedwards/scs/v2 v2.8.0
	github.com/felixge/httpsnoop v1.0.4
	github.com/gorilla/mux v1.8.1
	github.com/panyam/s3gen v0.0.6
	google.golang.org/grpc v1.64.0
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/adrg/frontmatter v0.2.0 // indirect
	github.com/alecthomas/chroma/v2 v2.14.0 // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/morrisxyang/xreflect v0.0.0-20231001053442-6df0df9858ba // indirect
	github.com/panyam/goutils v0.1.1 // indirect
	github.com/radovskyb/watcher v1.0.7 // indirect
	github.com/yuin/goldmark v1.7.1 // indirect
	github.com/yuin/goldmark-highlighting/v2 v2.0.0-20230729083705-37449abec8cc // indirect
	go.abhg.dev/goldmark/anchor v0.1.1 // indirect
	golang.org/x/sys v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
