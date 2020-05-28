FROM buildkite/puppeteer:latest
ARG DOCKER_FRONTEND
ARG DOCKER_BACKEND
ENV DOCKER_FRONTEND=$DOCKER_FRONTEND
ENV DOCKER_BACKEND=$DOCKER_BACKEND

RUN apt-get update && apt-get -y install netcat nginx

# We set up a local nginx instance to redirect to the frontend here.  This is 
# because square refuses redirect urls that it can't reach, so the website 
# would be broken when accessed only via a local machine name.  Square makes an 
# exception for "localhost" so we use this local reverse proxy as a workaround.
RUN printf "server {\n   listen 80;\n   listen [::]:80;\n   server_name localhost;\n   location / {\n      proxy_pass $DOCKER_FRONTEND;\n   }\n }\n" > /etc/nginx/sites-available/default

ENV FRONTEND=http://localhost

COPY package.json package-lock.json ./
RUN npm install
COPY . .
CMD until nc -z $(printf $DOCKER_BACKEND | sed 's~http[s]*://~~g') 80; do \
		echo "Waiting for backend ..."; \
		sleep 1; \
	done; \
	until nc -z $(printf $DOCKER_FRONTEND | sed 's~http[s]*://~~g') 80; do \
		echo "Waiting for frontend  ..."; \
		sleep 1; \
	done; \
	nginx && npx mocha . --reporter mocha-junit-reporter --reporter-options mochaFile=results.xml && nginx -s quit
