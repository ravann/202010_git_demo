set -x
docker stop go-test
docker rm go-test

docker run -d -p 8000:8000 --name go-test go-test-image
