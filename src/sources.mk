
GO_FILES = \
	main.go

GO_FILES_config = \
	config.go

GO_FILES_handlers = \
	fizzbuzz.go \
	fizzbuzz_test.go \
	stats.go \
	stats_test.go

GO_FILES_helpers = \
	parse_int.go \
	parse_int_test.go

GO_FILES_models = \
	fizzbuzz.go \
	stats.go

GO_FILES_router = \
	router.go

GO_FILES_services = \
	fizzbuzz.go \
	fizzbuzz_test.go \
	stats.go \
	stats_test.go

GO_FILES += $(foreach SRC_SUB_DIR,$(SRC_SUB_DIRS),$(addprefix $(SRC_SUB_DIR)/,$(GO_FILES_$(SRC_SUB_DIR))))
