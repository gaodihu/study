CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sv main.go
sudo docker build -t test-service .