package main

import "flag"

// Find flag
func SetFlag(str []string) ArgumentsType {
	res := ArgumentsType{command: str[0], path: str[1]}
	if !port.set {
		res.defaultPort()
	} else {
		res.port = port.value
	}
	return res
}

// With flag magic
type StringFlag struct {
	set   bool
	value string
}

func (sf *StringFlag) Set(x string) error {
	sf.value = x
	sf.set = true
	return nil
}

func (sf *StringFlag) String() string {
	return sf.value
}

var port StringFlag

func init() {
	flag.Var(&port, "port", "the port")
}
