FROM golang

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o /hw

FROM scratch

COPY --from=0 /hw /hw

CMD ["/hw"]