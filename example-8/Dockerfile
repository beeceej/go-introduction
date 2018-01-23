FROM golang:alpine as build-env

LABEL maintainer=beeceej
WORKDIR /go/src/app
ADD ./main.go main.go
RUN CGO_ENABLED=0 go-wrapper install -ldflags '-extldflags "-static"'
RUN ls -la

FROM scratch
WORKDIR /go/bin/app
COPY --from=build-env /go/bin/app /app
CMD [ "/app" ]
EXPOSE 8080