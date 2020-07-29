FROM golang AS builder
LABEL stage=builder
WORKDIR /go/src/github.com/vvatelot/airtable-slack-notify
COPY . .
RUN rm .env
RUN go get go/src/github.com/vvatelot/airtable-slack-notify && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o appexec .

FROM alpine AS final
ARG ENV
WORKDIR /
COPY --from=builder /go/src/github.com/vvatelot/airtable-slack-notify/appexec .
CMD [ "./appexec" ]
