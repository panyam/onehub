# STAGE 1: A container with pnpm and python3 is required
FROM node:18-alpine as npm_base

WORKDIR /app

COPY next.config.js package*.json /app

RUN npm install

# Finally, we run the NextJS app
CMD ["npm", "run", "dev"]
