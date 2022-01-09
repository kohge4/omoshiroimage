FROM golang:latest as build

WORKDIR /go/src/app

COPY go.* ./
RUN go mod download
# Copy local code to the container image.
COPY . ./
# Build the binary.
RUN go build /go/src/app/cmd/web/main.go
# Run the web service on container startup.


FROM chromedp/headless-shell:latest
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
COPY --from=build /go/src/app/main /tmp

CMD ["/tmp/main"]