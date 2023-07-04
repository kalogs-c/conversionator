PROJECT = turbosizenator

build-lib:
	go build -buildmode=c-shared -o C/$(PROJECT).so cmd/main.go

build-c: ./C/main.c ./C/turbosizenator.h
	gcc ./C/main.c -L. ./C/$(PROJECT) -o ./C/dist/$(PROJECT)
