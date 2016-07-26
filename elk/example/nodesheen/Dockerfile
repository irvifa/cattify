FROM node:argon

ADD src app/
WORKDIR app/
RUN npm install
EXPOSE 3000
cmd ["node","server.js"]
