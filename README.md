# Docker demos

Explore each directory in this repo for some demos on how Docker works!

_All commands are run from the top level directory of this repo_

## Basics

A minimal docker container that runs a bash script then exits.

Build your image tagging it "minimal"

    docker build -t minimal ./minimal

Now you can see your image

    docker images minimal

Run your image, this creates a new container with the name "hello"

    docker run --name hello minimal

Your container should print "Hello!" then exit. We can now look up that container by name

    docker ps --all --filter name=hello

And also quickly re-run that existing container

    docker start hello

We also can use `docker run` to create a new container that overrides the default command:

    docker run --name goodbye echo "Goodbye"

Both containers can be run again by name

    docker start hello
    docker start goodbye

If you want to run some task in a container, but don't want to save the container add `--rm`

    docker run --rm minimal ls /app


## Adding libraries

Build out image. This will install curl, Node.js, and our npm dependencies

    docker build -t node-demo ./dependencies

Let's run our new image, it should print out the status of our api

    docker run --name api-status node-demo


## Using our own base image

Starting with our Node.js image from the previous section, lets add our application

    docker build -t node-demo-app ./application

Let's start our application. Since it runs on port 8080, let's expose all ports with `-P`

    docker run --name hello-app -p 8080:8080 node-demo-app

We can stop the app from another terminal window

    docker kill hello-app


## Local application development

We want to use docker during development without having to rebuild between each change.

    docker build -t node-demo-app-dev ./application --file ./application/Dockerfile.dev

Since we're going to mount all of our code, let's install our deps on our machine

    cd application && npm install && cd ..

Let's mount our code as a volume so local changes will appear inside the container.

    docker run --name hello-app-dev -p 8080:8080 -v `pwd`/application/:/app node-demo-app-dev

Now make a change to the application and see the change take effect instantly!
