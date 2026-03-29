FROM node:25-alpine AS builder

WORKDIR /usr/src/app

COPY ./package.json ./package-lock.json ./

RUN npm ci --omit=dev

COPY . .

RUN npm run build

FROM node:25-alpine AS dev

WORKDIR /usr/src/app

COPY ./package.json ./package-lock.json ./

RUN npm i

COPY . .

EXPOSE 5173

CMD ["npm", "run", "dev"]

FROM nginx:alpine AS prod

COPY --from=builder /usr/src/app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]