#We wil use docker multistage build to redcue the image size as we don't need all the dependencies during the run time.
#We just need the binary file for the application to run. 

#Build stage
FROM golang:1.22.4-alpine3.20 as builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

#Run stage
#this alpine is linux apline not golang
FROM alpine:3.14 
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY app.env .
COPY start.sh .
COPY ./db/migration ./migration
RUN chmod +x start.sh

EXPOSE 8080
CMD [ "/app/main" ] 
ENTRYPOINT [ "/app/start.sh" ]