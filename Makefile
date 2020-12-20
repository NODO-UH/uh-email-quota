
build:
	go build -ldflags="-s -w" -o email-quota.bin -i src/main.go

compress: build
	upx --brute email-quota.bin
