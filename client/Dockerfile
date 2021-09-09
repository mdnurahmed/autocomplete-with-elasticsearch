# Dockerfile for React client

# Build react client
FROM node:10.16-alpine

# Working directory be app
WORKDIR /usr/src/app

COPY ./package.json ./
COPY ./package-lock.json ./

###  Installing dependencies

RUN yarn install 

# copy local files to app folder
COPY . .

EXPOSE 3000

CMD ["npm","start"]