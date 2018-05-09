CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cl main.go
sudo docker build -t test-cli .