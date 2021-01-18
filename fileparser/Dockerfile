FROM golang:1.15.6
WORKDIR /go/src/github.com/humansriot
RUN git clone https://github.com/humansriot/5-shades-of-isolation.git
WORKDIR /go/src/github.com/humansriot/5-shades-of-isolation/fileparser
RUN CGO_ENABLED=0 GOOS=linux go build

FROM debian:buster-slim  
RUN apt-get update && apt-get install -y ca-certificates x11-apps
WORKDIR /root/
COPY --from=0 /go/src/github.com/humansriot/5-shades-of-isolation/fileparser .
ENTRYPOINT ["./fileparser"] 