FROM golang:1.23

# https://github.com/oapi-codegen/oapi-codegen
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.2.0
