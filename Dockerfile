FROM golang:1.16.3
WORKDIR /go/src/github.com/kavemun/toy-robot/
COPY . .
RUN go test ./... -v -short -p 1 -cover
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o toy-robot .

FROM ubuntu:18.04
RUN apt-get update && apt-get -y install \
  autoconf \
  python \
  python-dev \
  python-distribute \
  python3-pip \
  unzip


# Install AWS CLI
RUN \
  apt-get -y install awscli

WORKDIR /app
COPY --from=0 /go/src/github.com/kavemun/toy-robot/toy-robot .
COPY ecs-worker.sh .

CMD [ "./ ecs-worker.sh" ]
#ENTRYPOINT [ "/app/toy-robot"]
