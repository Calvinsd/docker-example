# syntax=docker/dockerfile:1

FROM  node:16.14.0-alpine AS builder

LABEL maintainer="Calvin Dsouza"

ENV NODE_ENV=production

WORKDIR /app

COPY ["package.json","package-lock.json","./"]

RUN npm install --production

COPY . .

# Run the application
CMD ["node","./src/index.js"]