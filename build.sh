env GOOS=windows GOARCH=386 go build -o build/win/386/webclip.exe cmd/webclip.go
env GOOS=windows GOARCH=amd64 go build -o build/win/amd64/webclip.exe cmd/webclip.go
env GOOS=linux GOARCH=386 go build -o build/linux/386/webclip cmd/webclip.go
env GOOS=linux GOARCH=amd64 go build -o build/linux/amd64/webclip cmd/webclip.go
env GOOS=linux GOARCH=arm go build -o build/linux/arm/webclip cmd/webclip.go
env GOOS=linux GOARCH=arm64 go build -o build/linux/arm64/webclip cmd/webclip.go
