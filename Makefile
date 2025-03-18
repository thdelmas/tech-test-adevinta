# Fizz Buzz Makefile

NAME := fizzbuzz-api
EXEC_PATH := bin/$(NAME) 
RM = /bin/rm

.PHONY: all clean re

all: $(EXEC_PATH)

$(EXEC_PATH):
	@mkdir -p ./bin
	@go build -C src -o ../$(EXEC_PATH)

run: $(EXEC_PATH)
	@./$(EXEC_PATH)
