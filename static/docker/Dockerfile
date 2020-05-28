FROM jojomi/hugo:0.70.0 as hugobuilder

ARG HUGO_BASEURL
ARG HUGO_DYNAMIC

WORKDIR /app
COPY hugo .
RUN hugo

FROM node:14.3.0-alpine3.11 as nodebuilder

WORKDIR /app
COPY react/package.json react/package-lock.json ./
RUN npm install
COPY react .

RUN npx webpack --mode development

FROM nginx:1.17.10-alpine
COPY --from=hugobuilder /app/public /usr/share/nginx/html
COPY --from=nodebuilder /app/output /usr/share/nginx/html/js/react
