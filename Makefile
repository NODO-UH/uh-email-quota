# ======
# COLORS
# ======

NO_COLOR=\033[0m
RED_COLOR=\033[0;31m
BLUE_COLOR=\033[0;34m
YELLOW_COLOR=\033[1;33m

# =======
# SWAGGER
# =======

DEFINITIONS-PATH="swagger-spec/definitions"
PATHS-PATH="swagger-spec/paths"

swagger-mixin:
	@echo "$(YELLOW_COLOR)=====\nMIXIN\n=====$(NO_COLOR)"
	@echo "$(BLUE_COLOR)Mixin files: $(shell find $(DEFINITIONS-PATH) -type f -name "*.yml") $(shell find $(PATHS-PATH) -type f -name "*.yml") $(NO_COLOR)"
	swagger mixin swagger-spec/head.yml \
		$(shell find $(DEFINITIONS-PATH) -type f -name "*.yml") \
		$(shell find $(PATHS-PATH) -type f -name "*.yml") \
		-o swagger-spec/swagger.yml \
		--format=yaml

swagger-validate: swagger-mixin
	@echo "$(YELLOW_COLOR)========\nVALIDATE\n========$(NO_COLOR)"
	@echo "$(BLUE_COLOR)Validate file: swagger-spec/swagger.yml$(NO_COLOR)"
	swagger validate swagger-spec/swagger.yml

swagger-generate-server: swagger-validate
	@echo "$(YELLOW_COLOR)========\nGENERATE\n========$(NO_COLOR)"
	swagger generate server -A GestionEmailPlugin -f swagger-spec/swagger.yml

# =====
# BUILD
# =====

build: swagger-generate-server
	@echo "$(YELLOW_COLOR)=====\nBUILD\n=====$(NO_COLOR)"
	@echo "$(BLUE_COLOR)GOPATH=$(GOPATH)$(NO_COLOR)"
	go build -o plugin.bin -i cmd/gestion-email-plugin-server/main.go
	@echo "$(BLUE_COLOR)Binary: $(PWD)/gestion-email-plugin-server.bin$(NO_COLOR)"
	@echo "$(RED_COLOR)To start the server on a random port, run the command:$(BLUE_COLOR) ./plugin.bin$(NO_COLOR)"
	@echo "$(RED_COLOR)To start the server on 8080 port, run the command:$(BLUE_COLOR) ./plugin.bin --port 8080$(NO_COLOR)"
