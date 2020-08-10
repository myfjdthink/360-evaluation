FROM golang:alpine AS build

ENV GO111MODULE on
ENV GOPROXY https://goproxy.io,direct
RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base

WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-linkmode external -extldflags -static" -a -installsuffix -o 360-evaluation
RUN chmod +x 360-evaluation

FROM alpine as production
COPY --from=build /app/360-evaluation .
CMD ["./360-evaluation"]
