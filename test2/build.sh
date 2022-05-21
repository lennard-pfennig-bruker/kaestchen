set -e

export GOOS=linux
export GOARCH=386

go build -o main main.go

rm main.zip
zip main.zip main

openssl sha256 main.zip

aws s3 cp main.zip s3://go-test-lenne/

aws lambda update-function-code \
--function-name golangTest1 \
--region eu-central-1 \
--s3-bucket go-test-lenne \
--s3-key main.zip > lambda-update.log

curl -i https://op4wxb4za7vvbubufgdppxlksm0ojudm.lambda-url.eu-central-1.on.aws/