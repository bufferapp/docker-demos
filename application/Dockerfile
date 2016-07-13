# Let's use our previous container as a starting point
FROM node-demo

# Add our new application code
ADD . /app

# Install all of our application dependencies again
RUN npm install

# We want to tell our application to run on port 8080
ENV PORT 8080

# We tell Docker to expose port 8080 on the container to access out app
EXPOSE 8080

# Since our base image's default CMD is "node /app/index.js" we don't need to specify
# anything to override
