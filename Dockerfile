FROM golang:latest AS build
RUN mkdir /go/src/StringService_build
ADD ./src /go/src/StringService_build
WORKDIR /go/src/StringService_build
RUN go get -d ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main .

FROM scratch AS runtime
COPY --from=build /main /
EXPOSE 8080
CMD ["/main"]
