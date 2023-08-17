# STAGE 1: A container with pnpm and python3 is required
FROM node:18-alpine as npm_base

WORKDIR /app

COPY package*.json /app

RUN npm install

COPY . /app

# Finally, we run the NextJS app
CMD ["npm", "run", "dev"]
