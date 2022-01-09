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
WORKDIR /tmp
RUN apt-get update; apt install dumb-init -y
ENTRYPOINT ["dumb-init", "--"]
# マルチステージビルド https://matsuand.github.io/docs.docker.jp.onthefly/develop/develop-images/multistage-build/#use-an-external-image-as-a-stage
COPY --from=build /go/src/app/ /tmp/
#COPY --from=build /go/src/app/assets/ /tmp/assets

CMD ["/tmp/main"]