regenerate-docker:
	docker run --rm -it \
		-v `pwd`:/go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/messages \
		-w /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/messages \
		obraun/vss-protoactor-jenkins \
		make regenerate

regenerate:
	go install -v github.com/gogo/protobuf/protoc-gen-gogoslick
	protoc -I=. -I=${GOPATH}/src -I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogoslick_out=plugins=grpc:. messages.proto