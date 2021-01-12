# 5 shades of isolation

## the good old world

the operational world has never been without complications. Maybe you already know them, maybe not. The following exercise points out the issues that can be encountered in the deployment of an application, and to name a few:

- conflicts between applications
- not very clear instructions and insufficient documentation
- manual actions

## the challenge

the challenge is to run, by one means or another, these two applications in the same version on the same server, and to see if it is still possible.

## how to proceed

clone the repo:

`git clone https://github.com/humansriot/5-shades-of-isolation.git`

copy both binaries in your home directory

`cp 5-shades-of-isolation/fileparser/fileparser 5-shades-of-isolation/httptime/httptime .`

you can use the source code of the application and the documentation found under `5-shades-of-isolation/fileparser` and `5-shades-of-isolation/httptime` to troushoot both applications

### first challenge

just launch both applications:

`$ ./fileparser`

`$ ./httptime`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again

### second challenge

launch the two applications in version 1:

`$ ./fileparser -version 1`

`$ ./httptime -version 1`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again

### third challenge

launch the two applications in version 1:

`$ ./fileparser -version 2`

`$ ./httptime -version 2`

observe each other's behavior, resolve conflicts and find solutions so that it does not happen again
