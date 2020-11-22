#!/bin/bash
# chmod +x *.sh

docker pull mysql
sudo docker run -p 3310:3306 --name mysql_conn -v ~/mysql/test_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=112233  -d mysql

echo "waiting 45s"
sudo docker run -p 3310:3306 --name mysql_conn -v ~/mysql/test_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=112233  -d mysql
sudo docker cp inside.sh mysql_conn:/tmp/inside.sh
sleep 45
sudo docker exec -it mysql_conn /bin/sh /tmp/inside.sh