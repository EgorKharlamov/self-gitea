package main

func (obj *ArgumentsType) defaultPort() {
	if obj.port == "" {
		obj.port = "2221"
	}
}
