FROM ruby:2.3

ENV DEBIAN_FRONTEND noninteractive

RUN apt-get update && \
  apt-get install -y build-essential nodejs libpq-dev postgresql-client-9.4 --fix-missing --no-install-recommends
ADD src/ /app/
WORKDIR /app
RUN bundle install

EXPOSE 4114

CMD ["rails","s", "-p", "4114"]

