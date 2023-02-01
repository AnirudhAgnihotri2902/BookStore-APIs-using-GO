FROM golang:1.19-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="Anirudh Agnihotri <codewithanirudhagni@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy
COPY . .

EXPOSE 8080

#Command to run the executable
CMD ["go","run","start.go"]