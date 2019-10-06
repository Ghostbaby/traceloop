// syscalls.go based on traceleft's metagenerator:
// https://github.com/ShiftLeftSecurity/traceleft/blob/master/metagenerator/metagenerator.go
// Apache License 2.0

package straceback

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Param struct {
	Position int
	Name     string
}

type Syscall struct {
	Name   string
	Params []Param
}

var (
	syscallNames map[int]string
	cSyscalls    map[string]Syscall
)

func init() {
	err := gatherSyscalls()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("List of syscalls loaded.\n")
}

const syscallsPath = `/sys/kernel/debug/tracing/events/syscalls/`

// Converts a string to CamelCase
func toCamel(s string) string {
	s = strings.Trim(s, " ")
	n := ""
	capNext := true
	for _, v := range s {
		if v >= 'A' && v <= 'Z' || v >= '0' && v <= '9' {
			n += string(v)
		}
		if v >= 'a' && v <= 'z' {
			if capNext {
				n += strings.ToUpper(string(v))
			} else {
				n += string(v)
			}
		}
		if v == '_' || v == ' ' {
			capNext = true
		} else {
			capNext = false
		}
	}
	return n
}

var re = regexp.MustCompile(`\s+field:(?P<type>.*?) (?P<name>[a-z_0-9]+);.*`)

func parseLine(l string, idx int) (*Param, error) {
	n1 := re.SubexpNames()

	r := re.FindAllStringSubmatch(l, -1)
	if len(r) == 0 {
		return nil, nil
	}
	res := r[0]

	mp := map[string]string{}
	for i, n := range res {
		mp[n1[i]] = n
	}

	if _, ok := mp["type"]; !ok {
		return nil, nil
	}
	if _, ok := mp["name"]; !ok {
		return nil, nil
	}

	// ignore
	if mp["name"] == "__syscall_nr" {
		return nil, nil
	}

	var cParam Param
	cParam.Name = mp["name"]

	// The position is calculated based on the event format. The actual parameters
	// start from 8th index, hence we subtract that from idx to get position
	// of the parameter to the syscall
	cParam.Position = idx - 8

	return &cParam, nil
}

func parseSyscall(name, format string) (*Syscall, error) {
	syscallParts := strings.Split(format, "\n")
	var skipped bool

	var cParams []Param
	for idx, line := range syscallParts {
		if !skipped {
			if len(line) != 0 {
				continue
			} else {
				skipped = true
			}
		}
		cp, err := parseLine(line, idx)
		if err != nil {
			return nil, err
		}
		if cp != nil {
			cParams = append(cParams, *cp)
		}
	}

	return &Syscall{
		Name:   name,
		Params: cParams,
	}, nil
}

func gatherSyscalls() error {
	cSyscalls = make(map[string]Syscall)
	err := filepath.Walk(syscallsPath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == "syscalls" {
			return nil
		}

		if !f.IsDir() {
			return nil
		}

		eventName := f.Name()
		if strings.HasPrefix(eventName, "sys_exit") {
			return nil
		}

		syscallName := strings.TrimPrefix(eventName, "sys_enter_")

		formatFilePath := filepath.Join(syscallsPath, eventName, "format")
		formatFile, err := os.Open(formatFilePath)
		if err != nil {
			return nil
		}
		defer formatFile.Close()

		formatBytes, err := ioutil.ReadAll(formatFile)
		if err != nil {
			return err
		}

		cSyscall, err := parseSyscall(syscallName, string(formatBytes))
		if err != nil {
			return err
		}

		cSyscalls[cSyscall.Name] = *cSyscall

		return nil
	})
	if err != nil {
		return fmt.Errorf("error walking %q: %v", syscallsPath, err)
	}
	return nil
}

func syscallGetName(nr int) string {
	name, ok := syscallNames[nr]
	if !ok {
		return fmt.Sprintf("unknown(%d)", nr)
	}
	return name
}

func syscallGetCall(nr int, args [6]uint64, argsStr *[6]*string) string {
	name, ok := syscallNames[nr]
	if !ok {
		return fmt.Sprintf("unknown(%d)", nr)
	}

	ret := name + "("
	for i, p := range cSyscalls[name].Params {
		if i != 0 {
			ret += ", "
		}
		if i < 6 {
			if argsStr != nil && argsStr[i] != nil {
				ret += fmt.Sprintf("%q", *(*argsStr)[i])
			} else {
				ret += fmt.Sprintf("%v", args[i])
			}
		} else {
			ret += p.Name
		}
	}
	ret += ")"
	return ret
}

func syscallGetDef(nr int) (args [6]uint64) {
	if syscallNames[nr] == "execve" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "access" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "open" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "openat" {
		return [6]uint64{0, useNullByteLength, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "mkdir" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "chdir" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "pivot_root" {
		return [6]uint64{useNullByteLength, useNullByteLength, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "mount" {
		return [6]uint64{useNullByteLength, useNullByteLength, useNullByteLength, 0, 0, 0}
	}
	if syscallNames[nr] == "umount" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "umount2" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "sethostname" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "statfs" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "stat" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "lstat" {
		return [6]uint64{useNullByteLength, 0, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "newfstatat" {
		return [6]uint64{0, useNullByteLength, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "read" {
		return [6]uint64{0, useRetAsParamLength | paramProbeAtExitMask, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "write" {
		return [6]uint64{0, useArgIndexAsParamLength + 2, 0, 0, 0, 0}
	}
	if syscallNames[nr] == "getcwd" {
		return [6]uint64{useNullByteLength | paramProbeAtExitMask, 0, 0, 0, 0, 0}
	}
	return
}
