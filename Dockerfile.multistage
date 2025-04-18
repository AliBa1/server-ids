FROM golang:latest AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

FROM build-stage AS run-test-stage
RUN go test -v ./...

RUN CGO_ENABLED=1 GOOS=linux go build -o /server-ids

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /server-ids /server-ids

EXPOSE 8080

CMD ["/server-ids"]