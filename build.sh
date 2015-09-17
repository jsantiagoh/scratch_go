env GOOS=linux GOARCH=amd64 go build sandbox.go
docker build -t the_question .
docker run the_question
