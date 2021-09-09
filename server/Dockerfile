FROM golang:alpine AS build
WORKDIR /mybuild
COPY . .
# RUN go mod tidy
RUN GOOS=linux go build -ldflags="-s -w"


FROM alpine:3.13
WORKDIR /myapp
COPY app.env /myapp/
COPY --from=build /mybuild/autocomplete /myapp/
EXPOSE 8080
CMD ["./autocomplete"]