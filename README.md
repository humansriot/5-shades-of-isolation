# 5 shades of isolation

## part 1

- discussion on technology: why, pros, cons (15 min)
- setup instructions and handling (10 min)
- interlude (5 min)
- lab: [the good old world](## step 1: the good old world) (25 min)
- interlude (5 min)
- lab: [create a dockerfile, an image and then run it](## step 2: the world of containers) (25 min)
- interlude (5 min)
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

first try by copying the binary [Dockerfile best practices](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)

second try by building the binary [How to build and install go programs](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)

third try by building the binary and puting it in a new image [multistage build](https://docs.docker.com/develop/develop-images/multistage-build/)

then try with differents args, and watch it run [run containers](https://docs.docker.com/get-started/nodejs/run-containers/)

## part 2

- run a container with arguments: send logs to stdout (15 min)
- 10 minutes debriefing (10 min)
- interlude (5 min)
- run a container in Kubernetes (25 min)
- interlude (5 min)
- forget the arguments, and use a configuration file instead (20 min)
- 10 minutes debriefing (10 min)
- interlude (5 min)
- versioning our deployment with helm (20 min)
- 10 minutes debriefing (10 min)

## run a container with arguments: send logs to stdout

Maybe you noticed that `httptime` and` fileparser` had options on the command line:

```
$ ./httptime -h
Usage of ./httptime:
  -d string
    	the directory of static file to host (default ".")
  -p string
    	port to serve on (default "8100")
  -stdout
    	log to stdout instead of log file
  -version int
    	version of the app
```

You will find a ready-to-use Dockerfile in `5-shades-of-isolation/fileparser/` and `5-shades-of-isolationfileparser/httptime/`

### the challenge

Build the image, then run the container with the option `-stdout`

- How does the container behave when you try to pass it an argument?
- what modification could you make to Dockerfile to change this behavior? (help: https://phoenixnap.com/kb/docker-cmd-vs-entrypoint)

Stop for a moment to think about how to replace a container launched with the wrong option with another with the right option *without service interruption*.

is that even possible? what are the options?

## run a container in Kubernetes

Docker was not created to address availability issues, for that we need an orchestrator. since 2017, Kubernetes is the standard in terms of orchestration

### the challenge

Are you ready to deploy your first app in Kubernetes?

two deployment files are available to you, one for httptime `httptime/Kubernetes/deployment.yaml` and the other for fileparser `fileparser/Kubernetes/deployment.yaml`. what if you try to deploy these apps yourself? You will find some help [here](https://kubernetes.io/docs/concepts/cluster-administration/manage-deployment/)

as you can see, these files are "incomplete", because there is no argement `-sdtout` that went to containers. could you add it with the help of this documentation: [https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/](https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/)

## forget the arguments, and use a configuration file instead

in this exercise we will focus on `httptime`

as you can notice in the `Kubernetes` folder, a few resources have been added: all these resources are fairly standard to run an application in Kubernetes

### the challenge

There are two things wrong with this deployment, however:

1. the ops are complaining about not having access to the logs of the application through the command `kubectl logs $pod-name` (I know that you do not know the name of the pod, but I reassure you, me neither! but here is a memory aid [https://kubernetes.io/fr/docs/reference/kubectl/cheatsheet/](https://kubernetes.io/fr/docs/reference/kubectl/cheatsheet/))
1. customers are complaining that they cannot access the app ... can you find out why?

there is no one method to solve this problem, let your imagination run wild, we will share the solutions

## versionning our deployment with helm

as you have seen, modifying yaml files is not a piece of cake ...

does what is in your files reflect what is in the clutser? how are they versioned? and again, there is not a solution to this problem, but several solutions: Helm is one of them

let's continue working with `httptime`: in the` Kubernetes` directory, you have a *chart*

a chart is mainly made up of 3 components:

- `Chart.yaml` which allows you to define the chart version and this application
- `templates` is a directory containing all the files needed to deploy the application, but in the form of a Jinja-2 template
- `values.yaml` is the file which contains all the parameters necessary for the deployment

what is the purpose of creating a chart?

1. pack an application and the deployment configuration (which Docker cannot do for example)
1. only to interact with a small number of parameters, those contained in `Chart.yaml` and `values.yaml`
1. to know at all times the status of what is deployed in the Kubernetes cluster vs what is in the files

### the challenge

let's deploy `httptime` with Helm!

`$ helm install --generate-name ./httptime-chart`

oops! we made some mistakes:

1. we have configured the container to expose port 8080 instead of port 8100
1. we did not send the logs to `stdout`
1. we used image 1.0.0 instead of image 1.1.0

will you be able to correct these problems and deploy a new version of the application?

#### astuce

`$ helm list` will give you the name of your deployment

once you identify it and fix the issues, update this version with the command

`$ helm upgrade $deployment-name ./httptime-chart`
