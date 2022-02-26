package main

import (
	"fmt"
	"strings"
)

func isFuncDeclaration(line string) bool {
	if line == "" {
		return false
	}
	words := strings.Split(line, " ")
	if len(words) == 1 {
		return false
	}
	if !strings.ContainsRune(words[1], '(') {
		return false
	}
	return true
}

func extractFunc(lines []string) (string, string, error) {
	cpptext := strings.Join(lines, "\n") + "\n\n"
	hpptext := lines[0]
	i := strings.LastIndex(hpptext, ")")
	r := []rune(hpptext)
	r[i+1] = ';'
	return "\n" + string(r[:i+2]) + "\n", cpptext, nil
}

func isClassDeclaration(line string) bool {
	if line == "" {
		return false
	}
	words := strings.Split(line, " ")
	return words[0] == "class"
}

func extractClass(lines []string) (string, string, error) {
	hpptext := "\n" + lines[0] + "\n"
	cpptext := ""
	cname := strings.Split(lines[0], " ")[1]
	infoLogger.Printf("encountered class: %v\n", cname)
	lines = lines[1 : len(lines)-1]
	for i := 0; i < len(lines); i++ {
		l := strings.Trim(lines[i], "\r")
		if len(l) == 0 {
			continue
		}
		if l == "public:" || l == "private:" || l == "protected:" {
			hpptext += l + "\n"
			continue
		}
		if l[len(l)-1] == ';' {
			hpptext += l + "\n"
			continue
		}
		if l[len(l)-1] == '{' {
			mstart := i
			mend := i
			for ; strings.Trim(strings.Trim(lines[mend], "\r"), " ") != "}"; mend++ {
				if mend == len(lines)-1 {
					return "", "", fmt.Errorf("no closing bracket for method (fstart: %v)", mstart)
				}
			}
			infoLogger.Printf("extracted method for class %v: %v - %v\n", cname, mstart+1, mend+1)
			mlines := lines[mstart : mend+1]
			hppt, cppt, err := extractMethod(cname, mlines)
			if err != nil {
				return "", "", err
			}
			hpptext += hppt
			cpptext += cppt
			i = mend
		}
	}
	hpptext += "};\n"
	return hpptext, cpptext, nil
}

func extractMethod(cname string, lines []string) (string, string, error) {
	// hpptext := lines[0]
	// ci := strings.LastIndex(hpptext, "{")
	// r := []rune(hpptext)
	// r[ci] = ';'
	// hpptext = string(r[:ci+1]) + "\n"
	cpptext := ""
	words := strings.Split(lines[0], " ")
	dn := false
	line := ""
	for _, w := range words {
		if !dn && strings.ContainsRune(w, '(') {
			dn = true
			w = cname + "::" + w
		}
		line += w + " "
	}
	i := strings.LastIndex(line, ": ")
	if i == -1 {
		i = strings.LastIndex(line, "{")
	}
	cpptext += strings.Trim(line, " ") + "\n"
	hpptext := "\t" + strings.ReplaceAll(strings.Trim(line[:i], " "), cname+"::", "") + ";\n"
	for _, l := range lines[1 : len(lines)-1] {
		cpptext += "\t" + strings.Trim(l, " ") + "\n"
	}
	cpptext += "}\n\n"
	return hpptext, cpptext, nil
}

func isStructDeclaration(line string) bool {
	return strings.HasPrefix(line, "struct")
}

func extractStruct(lines []string) (string, string, error) {
	hpptext := ""
	cpptext := ""
	for _, l := range lines {
		hpptext += l + "\n"
	}
	return hpptext, cpptext, nil
}
