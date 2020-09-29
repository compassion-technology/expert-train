CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o humpty-dumpty-api
docker build -t timargheim/humpty-dumpty-api .
docker push timargheim/humpty-dumpty-api