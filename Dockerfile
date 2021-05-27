FROM golang:1.16
EXPOSE 8000-9000

WORKDIR /
COPY . .

RUN go build hello-port-range.go
RUN ls

CMD /hello-port-range -start=8000 -end=9000
