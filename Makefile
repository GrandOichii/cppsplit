testfile := test*
testspath := tests
dirs := $(shell ls $(testspath))
flags := $(testpath)/test.cpp -out tests/test1/out -log

all: build
	@exit

build:
	@echo Building app...
	@go build -o cppsplit

install:
	@go install

test: build
	@for f in $(dirs); do $(MAKE) -s testdir tp=${testspath}/$${f}; done
	@echo All tests run successfully!

testdir:
	${eval outpath := ${tp}/out}
	${eval sff := ${suffix ${shell ls ${tp}/main.c*}}}
	@echo Testing folder ${tp}
	@mkdir -p ${outpath}
	@rm -f ${outpath}/*
	@echo Running cppsplit 
	@./cppsplit ${tp}/${testfile} -out ${outpath} -log
	@echo Compiling out files...
	@if [ ${sff} = ".cpp" ]; then g++ $(tp)/main.cpp $(outpath)/*.cpp -o main; fi
	@if [ ${sff} = ".c" ]; then gcc $(tp)/main.c $(outpath)/*.c -o main; fi
	@echo Running out file...
	@./main
	@rm main
	