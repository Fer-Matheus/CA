FROM golang:1.20 AS build

WORKDIR /app

COPY  . ./

RUN go mod tidy

RUN go build -o /CA

FROM ubuntu

WORKDIR /app

COPY --from=build /CA /app/CA
COPY ./certs ./certs

EXPOSE 2302

# USER nonroot:nonroot

ENTRYPOINT [ "/app/CA" ]