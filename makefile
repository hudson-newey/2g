OUT_DIR=build

build:
	mkdir -p $(OUT_DIR)
	go build -o $(OUT_DIR)/2g ./src/main.go
	go build -o $(OUT_DIR)/2g-daemon ./daemon/main.go

.PHONY: clean

clean:
	rm -rf $(OUT_DIR)
