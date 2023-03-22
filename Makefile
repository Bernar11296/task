IMAGE_NAME=test
CONTAINER_NAME=test-task
PORT=8080

build:
	sudo docker build -t $(IMAGE_NAME) .
run-container:
	sudo docker run --rm -it --name $(CONTAINER_NAME) -p $(PORT):$(PORT) $(IMAGE_NAME)