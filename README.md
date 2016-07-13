# Docker demos

Explore each directory in this repo for some demos on how Docker works!

## Basics

A minimal docker container that runs a bash script then exits.

Build your image

    docker build -t minimal ./01-minimal

Now you can see your image

    docker images minimal

Run your image, this creates a new container

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

