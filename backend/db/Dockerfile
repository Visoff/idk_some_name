FROM node:20-alpine

WORKDIR /app/

COPY package*.json ./

RUN npm install

COPY . .

ENV POSTGRES_HOST=localhost
ENV POSTGRES_SSL=false
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=31415926
ENV POSTGRES_DATABASE=dev

ENV CORS_ORIGIN=*
ENV CORS_METHODS=*
ENV CORS_HEADERS=*

ENV APIKEY ""

EXPOSE 8080
CMD [ "npm", "run", "run" ]