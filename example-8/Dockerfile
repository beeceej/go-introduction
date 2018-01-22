FROM golang:latest as build-env
LABEL maintainer=beeceej
WORKDIR /
ADD ./main.go /main.go
RUN export CGO_ENABLED=0 && go build -o example-8 main.go

FROM alpine
WORKDIR /app
COPY --from=build-env example-8 .
RUN ls  -la
CMD ./example-8
EXPOSE 8080