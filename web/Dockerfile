FROM node:12-alpine as build-stage

WORKDIR /app
COPY ./ /app/
RUN npm ci
RUN npm run build

FROM nginx:alpine
COPY --from=build-stage /app/build/ /usr/share/nginx/html
COPY --from=build-stage /app/dev/nginx.conf /etc/nginx/conf.d/default.conf
