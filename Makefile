GOCMD=go
BINARY_NAME=imgconv
NUM_PARALLEL=4

build:
	$(GOCMD) build -o $(BINARY_NAME) -v

test:
	$(GOCMD) test -v -p $(NUM_PARALLEL) -coverprofile=cover.out .
	$(GOCMD) tool cover -html=cover.out -o cover.html

clean:
	$(GOCMD) clean
	rm -f $(BINARY_NAME)
	rm -f cover*
