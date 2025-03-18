# Fizz Buzz Makefile

NAME := fizzbuzz-api
EXEC_PATH := bin/$(NAME) 
RM = /bin/rm

SRC_DIR := ./src

.PHONY: all clean re

all: $(EXEC_PATH)

$(EXEC_PATH):
	@mkdir -p ./bin
	go -C $(SRC_DIR) build -o ../$(EXEC_PATH) ./

run: $(EXEC_PATH)
	@go -C $(SRC_DIR) run .

clean:
	$(RM) -rf bin

re: clean $(EXEC_PATH)

test: re
	go -C $(SRC_DIR) test -v --race ./...

docker_build: test
	docker build --no-cache -t $(NAME) .

deploy: docker_build
	docker run -p 8080:8080 $(NAME) 
