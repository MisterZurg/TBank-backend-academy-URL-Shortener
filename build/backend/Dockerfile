FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies
# and only redownloading them in subsequent builds if they change
COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download && go mod verify

# copy app logic
COPY ../../backend/cmd ./backend/cmd
COPY ../../backend/docs ./backend/docs
COPY ../../backend/config ./backend/config
COPY ../../backend/internal ./backend/internal
COPY ../../backend/urlerrors ./backend/urlerrors

RUN mkdir -p /usr/local/bin/
RUN go mod tidy
RUN go build -v -o /usr/local/bin/app ./backend/cmd

EXPOSE 1323

CMD ["app"]