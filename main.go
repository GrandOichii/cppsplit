package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	sourcePath string
	outPath string
	loggingEnabled bool

	infoLogger *log.Logger
	warningLogger *log.Logger
	errorLogger *log.Logger
)

func init() {
	flags := log.Llongfile | log.LUTC
	infoLogger = log.New(os.Stdout, "INFO:", flags)
	warningLogger = log.New(os.Stdout, "WARN:", flags)
	errorLogger = log.New(os.Stdout, "ERR:", flags)

	flag.StringVar(&outPath, "out", ".", "Out path")
	flag.BoolVar(&loggingEnabled, "log", false, "enable logging")
}

func exit() {
	flag.PrintDefaults()
	os.Exit(1)
}

func checkErr(err error) {
	if err != nil {
		errorLogger.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("err: cppsplit requires the path to the cpp file")
		exit()
	}
	sourcePath = os.Args[1]
	os.Args = os.Args[1:]
	flag.Parse()
	if !loggingEnabled {
		infoLogger.SetOutput(ioutil.Discard)
		warningLogger.SetOutput(ioutil.Discard)
		errorLogger.SetOutput(ioutil.Discard)
	}
	if len(outPath) == 0 {
		fmt.Println("err: out path can't be blank")
		exit()
	}
	infoLogger.Print("source directory path - " + sourcePath)
	infoLogger.Print("out directory path - " + outPath)
	cpptext, hpptext, err := splitCPPFile(sourcePath)
	checkErr(err)
	p := path.Join(outPath, strings.TrimSuffix(filepath.Base(sourcePath), filepath.Ext(sourcePath)))
	cppfn := p + ".cpp"
	hppfn := p + ".hpp"
	infoLogger.Printf("parsed lines, header file name: %v, cpp file name: %v", hppfn, cppfn)
	err = os.WriteFile(cppfn, []byte(cpptext), 0755)
	checkErr(err)
	err = os.WriteFile(hppfn, []byte(hpptext), 0755)
	checkErr(err)
}

func splitCPPFile(sourcePath string) (string, string, error) {
	cpptext := "#include \"" + filepath.Base(strings.TrimSuffix(filepath.Base(sourcePath), filepath.Ext(sourcePath)) + ".hpp") + "\"\n\n"
	hpptext := "#pragma once\n\n"
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return "", "", err
	}
	lines := strings.Split(string(data), "\n")
	// read all the header lines, transfer them to header file
	hli := -1
	for {
		hli++
		l := lines[hli]
		if strings.HasPrefix(l, "#include") {
			continue
		}
		if strings.HasPrefix(l, "using ") {
			continue
		}
		break
	}
	infoLogger.Printf("header lines: 0 - %v", hli)
	for i := 0; i < hli; i++ {
		hpptext += lines[i] + "\n"
	}
	i := hli
	for {
		if i >= len(lines) {
			break
		}
		l := lines[i]
		if isFuncDeclaration(l){
			// extract function declaration
			fstart := i
			fend := i
			for ; lines[fend] != "}\r" && lines[fend] != "}"; fend++ {
				if fend == len(lines)-1 {
					return "", "", fmt.Errorf("no closing bracket for function (fstart: %v)", fstart)
				}
			}
			infoLogger.Printf("extracted function: %v - %v\n", fstart+1, fend+1)
			flines := lines[fstart:fend+1]
			hppt, cppt, err := extractFunc(flines)
			if err != nil {
				return "", "", err
			}
			hpptext += hppt
			cpptext += cppt
			i = fend
		}
		if isClassDeclaration(l) {
			cstart := i
			cend := i
			for ; lines[cend] != "};\r" && lines[cend] != "};"; cend++ {
				if cend == len(lines)-1 {
					return "", "", fmt.Errorf("no closing bracket for class (cstart: %v)", cstart)
				}
			}
			infoLogger.Printf("extracted class: %v - %v\n", cstart+1, cend+1)
			clines := lines[cstart:cend+1]
			hppt, cppt, err := extractClass(clines)
			if err != nil {
				return "", "", err
			}
			hpptext += hppt
			cpptext += cppt
			i = cend
		}
		i++
	}
	return cpptext, hpptext, nil
}