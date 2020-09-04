#!/bin/bash
# chmod +x *.sh

echo "start mysql container ..."
docker pull mysql
sudo docker run -p 3314:3306 --name mysql_conn -v ~/mysql/test_data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=112233  -d mysql

echo "starting redis container ..."
docker pull redis
docker run -itd --name redis_conn -p 6380:6379 redis --requirepass "8888"
sudo docker cp inside.sh mysql_conn:/tmp/inside.sh
echo "waiting 45s"
sleep 45
sudo docker exec -it mysql_conn /bin/sh /tmp/inside.sh
echo "finish all steps"
