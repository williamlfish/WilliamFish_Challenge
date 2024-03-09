FROM golang:latest as builder
WORKDIR /build
COPY go.mod go.sum ./
COPY src ./src
COPY pkg ./pkg
COPY templates ./templates
RUN ls -lah ./
RUN GOOS=linux GOARCH=arm go build -o server ./src/...


FROM alpine
RUN adduser app --disabled-password
USER app
COPY --from=builder --chown=app:app /build/templates go test/templates
COPY --from=builder --chown=app:app /build/server /bin/server
CMD ["server"]