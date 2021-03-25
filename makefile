bld: 
	mkdir build && go build -o build/gitea ./src/main.go

clean:
	rm -r build
