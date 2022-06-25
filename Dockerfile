FROM golang as build
WORKDIR /app
COPY ["src", "go.mod", "go.sum", "./"]
RUN go install && go build

FROM gcr.io/distroless/base:debug as app
# FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=build /app/mysql-container .
ENTRYPOINT [ "./mysql-container"]
