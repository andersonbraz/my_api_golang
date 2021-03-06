FROM golang:alpine as builder
RUN mkdir /build 
ADD ./app/ /build/app/
##ADD ./cli/ /build/cli/
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN go get -d -v go.mongodb.org/mongo-driver/mongo \
    && go get -d -v gopkg.in/mgo.v2/bson \
    && go get -d -v gopkg.in/mgo.v2 \
    && go get -d -v github.com/gorilla/mux \
    && go get -d -v github.com/spf13/cobra \
    && go get -d -v github.com/spf13/viper \
    && go get -d -v github.com/mitchellh/go-homedir \
    && go get -d -v github.com/spf13/cobra 
## WORKDIR /build/cli/
RUN ls -laR /build/app/*
## RUN ./agent check
## RUN ./agent list
WORKDIR /build/app/
RUN ls -laR /build/app/*
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cleacgo -ldflags '-extldflags "-static"' -o main .

FROM scratch
EXPOSE 8030 
## WORKDIR /cli/
## COPY --from=builder /build/cli/ .
## RUN ./agent check
## RUN ./agent list
WORKDIR /app/
COPY --from=builder /build/app/main .
CMD ["./main"]