FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go env -w GOPROXY=https://goproxy.io,direct
RUN CGO_ENABLED=0 GOOS=linux  go build -o 360-evaluation

FROM scratch as production
WORKDIR /app
COPY --from=build /app/360-evaluation .
CMD ["/360-evaluation"]
