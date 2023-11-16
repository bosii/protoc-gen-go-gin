

.PHONY: gen install

gen:
	protoc -I . -I ./third_party --go-gin_out=. ./pb/*.proto

doc:
	protoc -I . -I ./third_party --openapi_out=. ./pb/*.proto