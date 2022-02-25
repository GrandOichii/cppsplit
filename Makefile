flags := tests/test1/test.cpp -out tests/test1/out -log

all: build -run
	@exit

build:
	@echo Building app...
	@go build

install:
	@go install

-run:
	@./cppsplit.exe $(flags)

test: build -run
	@echo Compiling out files...
	@g++ tests/test1/main.cpp tests/test1/out/*.cpp
	@./a.exe
	@rm a.exe
	@echo Tests ran successfully!