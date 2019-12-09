# Starting from a popular distro
FROM debian:jessie

# We'll need curl to install node
RUN apt-get update && apt-get -y install curl

# Let's install Node.js
RUN curl -sL https://deb.nodesource.com/setup_10.x | bash -
RUN apt-get install -y nodejs

# Add all of our application code
ADD . /app

# All commands after this point are run in the /app directory
WORKDIR /app

# Install all of our application dependencies
RUN npm install

# Run our application
CMD node ./index.js
