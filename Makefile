build:
	go build -o bin/templar main.go
run:
	go run main.go
test:
	go test ./...
release:
	GOOS=darwin go build -o bin/templar -ldflags="-s -w -X 'main.version=${RELEASE_VERSION}'" main.go
	zip -j bin/templar-macos-${RELEASE_VERSION}.zip bin/templar
	GOOS=windows go build -o bin/templar -ldflags="-s -w -X 'main.version=${RELEASE_VERSION}'" main.go
	zip -j bin/templar-windows-${RELEASE_VERSION}.zip bin/templar
	GOOS=linux go build -o bin/templar -ldflags="-s -w -X 'main.version=${RELEASE_VERSION}'" main.go
	zip -j bin/templar-linux-${RELEASE_VERSION}.zip bin/templar