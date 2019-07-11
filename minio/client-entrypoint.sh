#!/bin/sh
mc config host add archive ${MINIO_SERVER} ${MINIO_ACCESS_KEY} ${MINIO_SECRET_KEY}
cp /root/config.json /root/my.json
sed -i -e 's|<MINIO_ACCESS_KEY>|'"${MINIO_ACCESS_KEY}"'|g' /root/my.json
sed -i -e 's|<MINIO_SECRET_KEY>|'"${MINIO_SECRET_KEY}"'|g' /root/my.json
sed -i -e 's|<NATS_USER>|'"${NATS_USER}"'|g' /root/my.json
sed -i -e 's|<NATS_PWD_PLAIN>|'"${NATS_PWD_PLAIN}"'|g' /root/my.json
cat /root/my.json | mc admin config set archive
mc admin service restart archive