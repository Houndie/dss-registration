FROM node:14.3.0

WORKDIR /app
COPY gatsby/gatsby-config.js gatsby/package.json gatsby/package-lock.json gatsby/tsconfig.json ./
RUN npm install
CMD npx gatsby develop --host 0.0.0.0
