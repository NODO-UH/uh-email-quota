# ======
# COLORS
# ======

NO_COLOR=\033[0m
YELLOW_COLOR=\033[1;33m

# =====
# BUILD
# =====

build:
	@echo "$(YELLOW_COLOR)=====\nBUILD\n=====$(NO_COLOR)"
	go build -o plugin.bin -i src/main.go
