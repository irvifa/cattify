#!/bin/sh
set -e

cat config/filebeat.yml | sed "s/LOGSTASH_HOST/$LOGSTASH_HOST/" | sed "s/LOGSTASH_PORT/$LOGSTASH_PORT/" | sed "s/INDEX/$INDEX/" > filebeat.yml.tmp
cat filebeat.yml.tmp > filebeat.yml
rm filebeat.yml.tmp

exec "$@"