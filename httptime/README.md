# httptimme

httptime is a simple program to display the current time through the http protocol

## how to use it

by default, httptime listens on port 8100, and its logs are output in httptime.log in the current directory

`./httptime`

in addition, httptime is executed in version 0

## how to change the default behavior

### args

`-d`: to change the current directory
`-p`: to change the port the app is serving on
`-stdout`: to send logs to stdout
`-version x`: to change the application version

### config file

you can also create a config.yml:

```
# port to serve on
port: 8100

# the directory of static file to host
directory: /var/log

# log to stdout instead of log file
stdout: false
```

## issues and bugs

please notify contact@humansriot.com for any problem or security@humansriot.com for a security breach
