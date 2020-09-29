FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY humpty-dumpty-api humpty-dumpty-api
ENTRYPOINT [ "./humpty-dumpty-api" ]
EXPOSE 80