module protoc-gen-go-gin

go 1.20

require (
	google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17
	google.golang.org/protobuf v1.31.0
)

require google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect

replace google.golang.org/protobuf => ./protobuf
