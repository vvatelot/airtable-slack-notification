FROM golang AS builder
LABEL stage=builder
WORKDIR /go/src/github.com/dktunited/ddfs-checker
COPY . .
RUN rm .env
RUN go get github.com/dktunited/ddfs-checker \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o appexec .

FROM alpine AS final
ARG ENV
WORKDIR /
COPY --from=builder /go/src/github.com/dktunited/ddfs-checker/appexec .
CMD [ "./appexec" ]
