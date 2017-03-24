# Server cross compile
```
env GOOS=linux GOARCH=386 go build -v iconthin
env GOOS=windows GOARCH=386 go build -v iconthin
```