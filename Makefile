.PHONY clean:
clean:
	rm main

.PHONY build:
build:
	go build main.go

.PHONY run:
run:
	./main
