package main

import (
	"os"
	"path/filepath"
	"strings"
)

func parse(s string) {
	name := s

	if O.OnlySuffix {
		suffix := filepath.Ext(name)
		if len(suffix) > 0 {
			pr("%s", suffix[1:])
		}
		return
	}

	if O.Path {
		name = filepath.Base(name)
	}

	if O.Suffix {
		suffix := filepath.Ext(s)
		if len(suffix) > 0 {
			name = strings.TrimSuffix(name, suffix)
		}
	}

	if O.Name {
		name = filepath.Dir(name)
	}

	if O.Absolute {
		var err error
		name, err = filepath.Abs(name)
		if err != nil {
			epr("Error: %s", err.Error())
			os.Exit(2)
		}
	}

	pr("%s", name)
}
