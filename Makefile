.PHONY: build
build: 
	go build -o bin/terraform-provider-oceanbase ./main.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -rf bin/*
