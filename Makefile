GOCMD=go
BINARY_NAME=imgconv

build:
	$(GOCMD) build -o $(BINARY_NAME) -v

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)
