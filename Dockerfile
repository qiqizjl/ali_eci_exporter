
#
# build container
#
FROM golang:1.12-alpine as builder
#WORKDIR /go/src/github.com/oliver006/redis_exporter/
WORKDIR /work
ADD ./ /work

RUN apk --no-cache add ca-certificates git
RUN BUILD_DATE=$(date +%F-%T) && CGO_ENABLED=0 GOOS=linux go build -o /ali_eci_exporter -mod=vendor


#
# Alpine release container
#
FROM scratch as scratch

COPY --from=builder /ali_eci_exporter /ali_eci_exporter
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE     8080
ENTRYPOINT [ "/ali_eci_exporter" ]
