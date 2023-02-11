FROM golang:1.20.0-alpine3.17 AS builder

ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=auto
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

WORKDIR /go/src
RUN echo $GOPATH

COPY . .
COPY . .

RUN apk -U add ca-certificates
ARG PROJECT_NAME=order
RUN echo ${PROJECT_NAME}
RUN go mod download

RUN go build ${PROJECT_NAME}/main.go

FROM scratch AS runner

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/main /

ENTRYPOINT ["./main"]