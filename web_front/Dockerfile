FROM node:18.18.0

WORKDIR /app

COPY package.json package-lock.json ./

COPY . .

RUN yarn install

EXPOSE 8080

CMD ["yarn","dev"]
