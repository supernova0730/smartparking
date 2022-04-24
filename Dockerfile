# syntax=docker/dockerfile:1

FROM golang AS build

ENV GOPATH /root/go/
ENV PROJECT ${GOPATH}/src/smartparking/
ENV BIN ${GOPATH}/bin/smartparking

WORKDIR ${PROJECT}
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o ${BIN}


FROM gcr.io/distroless/base-debian10

WORKDIR /root/
COPY --from=build /root/go/bin/smartparking smartparking
COPY config.yml .

EXPOSE 8080
CMD [ "/root/smartparking", "start", "--config=config.yml" ]