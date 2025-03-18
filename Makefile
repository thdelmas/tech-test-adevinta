# Fizz Buzz Makefile

NAME := fizzbuzz-api
EXEC_PATH := bin/$(NAME) 
RM = /bin/rm

SRC_DIR := ./src
SRC_SUB_DIRS := \
	config \
	handlers \
	helpers \
	models \
	router \
	services

include $(SRC_DIR)/sources.mk

SRC_FILES = $(addprefix $(SRC_DIR)/,$(GO_FILES))
SRC_FILES += $(SRC_DIR)/go.mod
SRC_FILES += $(SRC_DIR)/go.sum
SRC_FILES += $(SRC_DIR)/sources.mk

SRC_SUB_PATHS := $(addprefix ./, $(SRC_SUB_DIRS))

.PHONY: all clean re

all: $(EXEC_PATH)

$(EXEC_PATH): $(SRC_FILES)
	@mkdir -p ./bin
	go -C $(SRC_DIR) build -o ../$(EXEC_PATH)

run: $(EXEC_PATH)
	@./$(EXEC_PATH)

clean:
	$(RM) -rf bin

re: clean $(EXEC_PATH)

test: re
	go -C $(SRC_DIR) test -v $(SRC_SUB_PATHS)

docker_build: test
	docker build --no-cache -t $(NAME) .

deploy: docker_build
	docker run -p 8080:8080 $(NAME) 
