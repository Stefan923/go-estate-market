FROM node:20.11.0 as build

WORKDIR /usr/local/app
COPY ./ /usr/local/app/

RUN npm install -g @angular/cli
RUN npm install

EXPOSE 4200

CMD ["ng", "serve", "--host", "0.0.0.0"]
