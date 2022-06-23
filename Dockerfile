FROM golang as build
WORKDIR /go/src/app
COPY ["src", "go.mod", "go.sum", "./"]
RUN go install
CMD ["go", "run", "main.go"]


# FROM gcr.io/distroless/base:debug
# # FROM gcr.io/distroless/base
# WORKDIR /app
# COPY --from=build /go/bin/mysql-container /app
# CMD ["./mysql-container"]
