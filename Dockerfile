FROM golang:1.18

ARG SERVER_PORT

WORKDIR $GOPATH/src/paróquia/Scheduler

COPY . .

RUN go get -d -v ./...

# RUN go install -v ./...

RUN go build

EXPOSE ${SERVER_PORT}

CMD ["Schedule"]
