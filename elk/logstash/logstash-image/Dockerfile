FROM logstash:2.0

ENV DEBIAN_FRONTEND noninteractive

ADD sample.conf /etc/logstash/conf.d/sample.conf

EXPOSE 11000
CMD ["logstash","-f","/etc/logstash/conf.d/sample.conf"]

