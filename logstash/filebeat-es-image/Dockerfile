FROM      ubuntu:14.04

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update
RUN apt-get install -y curl 
ADD filebeat_install.sh /build/filebeat_install.sh
RUN /build/filebeat_install.sh
COPY config/filebeat.yml /filebeat.yml 
ADD config /config
ADD init.sh /init.sh
ENTRYPOINT ["/init.sh"]
CMD ["/usr/bin/filebeat", "-c", "/filebeat.yml"]
