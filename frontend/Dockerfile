FROM node:18.6-alpine3.16

WORKDIR /app/ebiznes_react

COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]