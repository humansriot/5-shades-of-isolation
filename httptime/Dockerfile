FROM golang:1.15.6
WORKDIR /go/src/github.com/humansriot
RUN git clone https://github.com/humansriot/5-shades-of-isolation.git
WORKDIR /go/src/github.com/humansriot/5-shades-of-isolation/httptime
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/humansriot/5-shades-of-isolation/httptime .
CMD ["./httptime"] 