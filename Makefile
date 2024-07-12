install:
	mkdir -p ~/.steampipe/plugins/local/diodb/
	go build -o  ~/.steampipe/plugins/local/diodb/diodb.plugin *.go
	mkdir -p ~/.steampipe/config/
	cp diodb.spc ~/.steampipe/config/

clean:
	rm -rf ~/.steampipe/plugins/local/diodb/
	rm ~/.steampipe/config/diodb.spc
