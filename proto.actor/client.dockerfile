FROM obraun/vss-protoactor-jenkins as builder
COPY . /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor
WORKDIR /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor
RUN go build -o client/client client/client.go

FROM iron/go
COPY --from=builder /go/src/github.com/ob-vss-ss19/ob-vss-ss19/proto.actor/client/client /app/client
EXPOSE 8090
ENTRYPOINT ["/app/client"]
CMD ["--bind=client:8090", "--remote=server:8091"]