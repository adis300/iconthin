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

# Install Postgres
http://imperialwicket.com/aws-install-postgresql-on-amazon-linux-quick-and-dirty/
```
sudo yum install postgresql postgresql-server postgresql-devel postgresql-contrib postgresql-docs
sudo service postgresql initdb
sudo vim /var/lib/pgsql9/data/postgresql.conf
Uncomment line 59:
#listen_addresses = 'localhost'
And update the line to enable connections from any IP address:
listen_addresses='*'
And uncomment line 63:
#port = 5432
sudo service postgresql start
```
# Login to Postgres
```
sudo su - postgres
psql -U postgres

CREATE USER iconthin WITH PASSWORD '$password';
CREATE DATABASE $dbname;
```