# docker-network-go-mysql

This project demonstrated the running of two docker containers on the same network and interacting with each other.

Firstly, run the mysql container using the below command which provides the relevant environment variables to the docker container as well.

```
docker run -d --name mysql-container --network mynetwork -e MYSQL_ROOT_PASSWORD=rishavkumar -e MYSQL_DATABASE=demodb -p 3306:3306 mysql:latest
```

Now, build the go-app image using the below command.

```
docker build -t go-app-image .
```

Finally, run the go-app container using the below command which provides the relevant environment variables as well to the docker container.

```
docker run -d --name go-app --network mynetwork -p 8500:8500 -e DB_USER=root -e DB_PASSWORD=rishavkumar -e DB_HOST=mysql-container -e DB_PORT=3306 -e DB_NAME=demodb go-app-image:latest
```

After hitting the /add post api, you can check the mysql db for the data to confirm that its working properly by using the below commands.

```
docker exec -it <mysql-container-id> bash
mysql -u root -p       <!-- password is rishavkumar -->
show databases;        <!-- shows the list of databases -->
use demodb;            <!-- demodb becomes the current db -->
select database();     <!-- shows the current selected database -->
show tables;           <!-- shows the list of tables in the current database -->
select * from users;   <!-- shows the list of rows in users table -->
```