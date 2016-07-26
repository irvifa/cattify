FROM node:argon

ADD src app/
WORKDIR app/
RUN npm install
EXPOSE 15120
cmd ["node","server.js"]
