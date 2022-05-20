set -e

export GOOS=linux
export GOARCH=x86_64

go build -o main main.go
zip main.zip main
openssl sha256 main.zip
aws s3 cp main.zip s3://go-test-lenne/