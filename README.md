# Server cross compile
```
env GOOS=linux GOARCH=386 go build -v iconthin
env GOOS=windows GOARCH=386 go build -v iconthin
```

# Connect
```
chmod 400 iconthin-key.pem
ssh -i "iconthin-key.pem" ec2-user@52.60.247.66
```
# Deploy to instance
```
scp -i iconthin-key.pem release.zip ec2-user@52.60.247.66:~/iconthin/release.zip
```

# Kill application on port
sudo lsof -n -i :80 | grep LISTEN