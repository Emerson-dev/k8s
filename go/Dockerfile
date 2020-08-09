FROM golang:1.11 as builder

ENV GO111MODULE=on

WORKDIR /src

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix "static" k8s

#ENTRYPOINT ["/go/bin/k8s"]
#ENTRYPOINT ["k8s"]

FROM scratch

USER 1000

COPY --from=builder /go/bin/k8s /bin/k8s

EXPOSE 8000

ENTRYPOINT ["/bin/k8s"]
