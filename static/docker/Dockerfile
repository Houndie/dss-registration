FROM node:14.3.0

WORKDIR /app
COPY gatsby-config.js package.json package-lock.json tsconfig.json gatsby-browser.js gatsby-node.js ./
RUN npm install
CMD npx gatsby develop --host 0.0.0.0
