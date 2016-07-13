from ruby:2.3

MAINTAINER Irvi Firqotul Aini <irvi.fa@gmail.com>
ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
  apt-get install -y build-essential nodejs libpq-dev postgresql-client-9.4 --fix-missing --no-install-recommends
ADD Gemfile /app/
ADD cattify/ /app/
WORKDIR /app
RUN bundle install
WORKDIR cattify

EXPOSE 3000
CMD ["rails","s"]
