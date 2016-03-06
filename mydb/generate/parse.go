package main

import (
	"errors"
	"fmt"
	"strings"
)

type StructData struct {
	PackName  string
	IntfPack  string   // Tag package definition with FOR xxxx
	Imports   []string // include intf pack if needed
	Name      string
	LowerName string      // sql tablename, basically
	Fields    []FieldData // Tag fields with // comments
	PKStr     string
	Writeable bool
}

type FieldData struct {
	Name      string
	LowerName string // sql colname, json name
	Type      string // include import path if needed
	SQLType   string
	// sql
	PK         bool // tag PK for true: order pk fields in struct
	Update     bool // considered default: tag NOUPDATE for false
	Insert     bool // considered default: tag NOINSERT for false
	InsertScan bool // tag SCAN for true
	Comparable bool // tag NOCOMP for false
	CanNull    bool // tag NULL for true
}

func ParseStructData(inBytes []byte) (StructData, error) {
	sD := StructData{}
	inStr := string(inBytes)
	inLines := strings.Split(inStr, "\n")
	// find package name
	var start int
	for i, line := range inLines {
		if err := BadLine(line); err != nil {
			return sD, err
		}
		var comments []string
		line, comments = ScrubLine(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "package ") {
			start = i
			parts := strings.Fields(line)
			sD.PackName = parts[1]
			for j, c := range comments {
				if j == len(comments)-1 {
					break
				}
				if c == "FOR" {
					sD.IntfPack = comments[j+1] + "."
					break
				}
			}
		} else {
			return sD, fmt.Errorf("line %d: file did not start with package declaration", i)
		}
		break
	}
	if sD.PackName == "" {
		return sD, errors.New("no package name found")
	}
	for i, line := range inLines[start:] {
		if err := BadLine(line); err != nil {
			return sD, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "import") {
			if err := (&sD).ImportParse(inLines[start+i:]); err != nil {
				return sD, err
			}
		} else if strings.HasPrefix(line, "type") {
			if err := (&sD).TypeParse(inLines[start+i:]); err != nil {
				return sD, err
			}
			return sD, nil
		}
	}
	return sD, errors.New("no struct data found")
}

func BadLine(line string) error {
	if strings.Contains(line, "/*") || strings.Contains(line, "*/") {
		return errors.New("mulegen does not support multi-line comments")
	}
	return nil
}

func ScrubLine(line string) (string, []string) {
	parts := strings.Split(line, "//")
	line = strings.TrimSpace(parts[0])
	if len(parts) == 1 {
		return line, nil
	}
	comments := strings.Fields(strings.Join(parts[1:], ""))
	return line, comments
}

func ParseFieldComments(comments []string) map[string]bool {
	mp := make(map[string]bool, len(comments))
	for _, c := range comments {
		mp[strings.ToUpper(c)] = true
	}
	return mp
}

func (sd *StructData) ImportParse(lines []string) error {
	fields := strings.Fields(lines[0])
	if fields[1] != "(" {
		sd.Imports = []string{strings.Trim(fields[1], "\"")}
		return nil
	}
	for i, line := range lines {
		if i == 0 {
			continue
		}
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ")") {
			return nil
		}
		if strings.HasPrefix(line, "type") {
			return errors.New("type declaration found in import parsing")
		}
		parts := strings.Split(line, "\"")
		if len(parts) < 2 {
			continue
		}
		sd.Imports = append(sd.Imports, parts[1])
	}
	return errors.New("no import ending found")
}

func (sd *StructData) TypeParse(lines []string) error {
	line, comments := ScrubLine(lines[0])
	fields := strings.Fields(line)
	cMap := ParseFieldComments(comments)
	if !cMap["NOWRITE"] {
		sd.Writeable = true
	}
	if len(fields) < 3 {
		return errors.New("couldn't parse first struct line")
	}
	if fields[2] != "struct" {
		return errors.New("non-struct type declared")
	}
	sd.Name = fields[1]
	sd.LowerName = strings.ToLower(sd.Name)
	var pks []string
	for i, line := range lines {
		if i == 0 {
			continue
		}
		line, comments = ScrubLine(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "}") {
			sd.PKStr = strings.Join(pks, ", ")
			return nil
		}
		fD := FieldData{}
		fields = strings.Fields(line)
		if len(fields) < 2 {
			return fmt.Errorf("no type found for field %s", fields[0])
		}
		fD.Name = fields[0]
		fD.LowerName = strings.ToLower(fD.Name)
		fD.Type = fields[1]

		cMap := ParseFieldComments(comments)
		if cMap["PK"] {
			fD.PK = true
			pks = append(pks, fD.LowerName)
		}
		if cMap["SCAN"] {
			fD.InsertScan = true
		}
		if !cMap["NOINSERT"] {
			fD.Insert = true
		}
		if !cMap["NOUPDATE"] {
			fD.Update = true
		}
		if !cMap["NOCOMP"] {
			fD.Comparable = true
		}
		if cMap["NULL"] {
			fD.CanNull = true
		}
		switch fields[1] {
		case "db.IntList":
			fD.SQLType = "int[]"
			fD.Comparable = false
		case "sql.NullInt64":
			fD.SQLType = "int"
			fD.CanNull = true
		case "string":
			fD.SQLType = "text"
		case "db.StringList":
			fD.SQLType = "text[]"
			fD.Comparable = false
		case "sql.NullString":
			fD.SQLType = "text"
			fD.CanNull = true
		case "sql.NullBool":
			fD.SQLType = "bool"
			fD.CanNull = true
		case "hexagon.Coord":
			fD.SQLType = "point"
		case "hexagon.NullCoord":
			fD.SQLType = "point"
			fD.CanNull = true
		case "hexagon.CoordList":
			fD.SQLType = "point[]"
			fD.Comparable = false
		default:
			fD.SQLType = fields[1]
		}
		sd.Fields = append(sd.Fields, fD)
	}
	return errors.New("no type declaration ending found")
}
