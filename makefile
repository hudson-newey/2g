OUT_DIR=build

.PHONY: clean
.PHONY: docs

build:
	mkdir -p $(OUT_DIR)
	go build -o $(OUT_DIR)/2g ./src/main.go
	go build -o $(OUT_DIR)/2g-daemon ./daemon/main.go

clean:
	rm -rf $(OUT_DIR)

docs:
	pandoc --standalone --to man ./README.md -o ./2g.1
