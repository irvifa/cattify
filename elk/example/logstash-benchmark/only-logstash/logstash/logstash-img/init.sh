#!/bin/sh
set -e

exec /opt/logstash/bin/logstash -f /etc/conf.d/sample.conf  >> /opt/logs/logstash.log 2>&1

