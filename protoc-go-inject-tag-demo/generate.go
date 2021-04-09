package generate

//go:generate kratos proto client .
//go:generate go get github.com/favadi/protoc-go-inject-tag
//go:generate sh protoc-go-inject-tag.sh
