FROM ruby:2.3

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get install -y wget apt-transport-https ca-certificates openjdk-7-jre
RUN wget -qO - https://packages.elastic.co/GPG-KEY-elasticsearch | apt-key add -
RUN echo "deb https://packages.elastic.co/logstash/2.3/debian stable main" | tee -a /etc/apt/sources.list
RUN apt-get update 
RUN apt-get install -y logstash curl

ADD logstash.conf /etc/logstash/conf.d/logstash.conf

VOLUME ["/opt/logstash/bin","/etc/logstash/conf.d","/opt/logs"]
EXPOSE 11001
CMD ["/opt/logstash/bin/logstash","-f","/etc/logstash/conf.d/logstash.conf"]
