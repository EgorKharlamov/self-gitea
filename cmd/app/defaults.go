package main

// Paste default port if not exists
func (obj *ArgumentsType) defaultPort() {
	if obj.port == "" {
		obj.port = "2221"
	}
}
