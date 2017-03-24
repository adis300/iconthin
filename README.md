# Server cross compile
```
env GOOS=linux GOARCH=386 go build -v iconthin
env GOOS=windows GOARCH=386 go build -v iconthin
```

# Connect
ssh -i "coturn-key.pem" ec2-user@ec2-52-10-224-174.us-west-2.compute.amazonaws.com
