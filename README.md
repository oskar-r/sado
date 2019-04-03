###Install
docker-compose -up --build

You need to inject new config to minio to enable events
mc config host add mys3 http://localhost:9001 oskar z2yByK2hB1ssIdddJtt3uql@l2gx
cat minio/config.json | mc admin config set docker &&  mc admin service restart docker

