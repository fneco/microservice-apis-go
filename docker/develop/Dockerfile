FROM golang:1.23

RUN <<EOF

go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0

go install github.com/air-verse/air@v1.60.0

EOF

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
