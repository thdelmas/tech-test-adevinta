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


.PHONY: all clean re

all: $(EXEC_PATH)

$(EXEC_PATH): $(SRC_FILES)
	@mkdir -p ./bin
	go build -C src -o ../$(EXEC_PATH)

run: $(EXEC_PATH)
	@./$(EXEC_PATH)
