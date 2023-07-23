build:
	go build -o bin/main main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run