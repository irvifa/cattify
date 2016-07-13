from ruby:2.3

MAINTAINER Irvi Firqotul Aini <irvi.fa@gmail.com>
ENV DEBIAN_FRONTEND noninteractive

ADD Gemfile /app/
ADD index.rb /app/
WORKDIR /app
RUN bundle install

EXPOSE 4567
CMD ["ruby","index.rb"]
