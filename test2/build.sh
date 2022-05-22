set -e

export GOOS=linux
export GOARCH=386

go build -o main main.go

rm deployment.zip || true
zip deployment.zip main

aws s3 cp deployment.zip s3://go-test-lenne/

aws lambda update-function-code \
--function-name golangTest1 \
--region eu-central-1 \
--s3-bucket go-test-lenne \
--s3-key deployment.zip > lambda-update.log

curl -i https://op4wxb4za7vvbubufgdppxlksm0ojudm.lambda-url.eu-central-1.on.aws/