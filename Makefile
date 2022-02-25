flags := tests/test1/test.cpp -out tests/test1/out -log

all: build -run
	@exit

build:
	@echo Building app...
	@go build -o cppsplit

install:
	@go install

-run:
	@./cppsplit $(flags)

test: build -run
	@echo Compiling out files...
	@g++ tests/test1/main.cpp tests/test1/out/*.cpp -o main
	@./main
	@rm main
	@echo Tests ran successfully!
