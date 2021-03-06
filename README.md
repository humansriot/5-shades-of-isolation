# 5 shades of isolation

## part 1

- discussion on technology: why, pros, cons (15 min)
- setup instructions and handling (10 min)
- break (5 min)
- lab: [the good old world](## step 1: the good old world) (25 min)
- break (5 min)
- lab: [create a dockerfile, an image and then run it](## step 2: the world of containers) (25 min)
- break (5 min)
- debrief on what we have learned so far and discussions (25 min)

## step 1: the good old world

the operational world has never been without complications. Maybe you already know them, maybe not. The following exercise points out the issues that can be encountered in the deployment of an application, and to name a few:

- conflicts between applications
- not very clear instructions and insufficient documentation
- manual actions

### the challenge

the challenge is to run, by one means or another, these two applications in the same version on the same server, and to see if it is still possible.

### how to proceed

clone the repo:

`git clone https://github.com/humansriot/5-shades-of-isolation.git`

copy both binaries in your home directory

`cp 5-shades-of-isolation/fileparser/fileparser 5-shades-of-isolation/httptime/httptime .`

you can use the source code of the application and the documentation found under `5-shades-of-isolation/fileparser` and `5-shades-of-isolation/httptime` to troushoot both applications

#### first challenge

just launch both applications:

`$ ./fileparser`

`$ ./httptime`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again

#### second challenge

launch the two applications in version 1:

`$ ./fileparser -version 1`

`$ ./httptime -version 1`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again

#### third challenge

launch the two applications in version 2:

`$ ./fileparser -version 2`

`$ ./httptime -version 2`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again

## step 2: the world of containers

### goals of containers

the containers came to solve these isolation problems. For 50 years containers have been fixing isolation problems, so what has changed with Docker?

Docker has made the process of creating images and running them super easy

### the challenge

the challenge is to run the applications on the same host in containers

#### lecture

FROM instruction in Dockerfile

#### first challenge: create a dockerfile

1. first try by copying the binary [Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
1. second try by building the binary [How to build and install go programs](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)
1. third try by building the binary and puting it in a new image [multistage build](https://docs.docker.com/develop/develop-images/multistage-build/)
1. then try with differents args, and watch it run [run containers](https://docs.docker.com/get-started/nodejs/run-containers/)
