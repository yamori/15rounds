buildrun:
	go build -o bin/15rounds_main main.go
	./bin/15rounds_main

clean:
	rm bin/*