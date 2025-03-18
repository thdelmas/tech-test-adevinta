
GO_FILES = \
	main.go

GO_FILES_config = \
	config.go

GO_FILES_handlers = \
	fizzbuzz.go \
	stats.go

GO_FILES_helpers = \
	parse_int.go

GO_FILES_models = \
	models.go

GO_FILES_router = \
	router.go

GO_FILES_services = \
	fizzbuzz.go \
	stats.go

GO_FILES += $(foreach SRC_SUB_DIR,$(SRC_SUB_DIRS),$(addprefix $(SRC_SUB_DIR)/,$(GO_FILES_$(SRC_SUB_DIR))))
