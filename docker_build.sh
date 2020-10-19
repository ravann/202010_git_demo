docker stop go-test
docker rm go-test
docker rmi go-test-image

docker build --rm -f Dockerfile -t go-test-image:latest .
