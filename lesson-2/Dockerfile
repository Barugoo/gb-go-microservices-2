#1 
# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/myapp .

# FROM scratch
# COPY bin/myapp /
# CMD ["/myapp"]


# #2 Compilation on container build
# FROM golang:1.15

# ADD . /myapp
# WORKDIR /myapp

# # Собираем бинарный файл
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
#    go build -o /myapp .

# CMD ["./myapp"]





# #3 Multi-Stage

# FROM golang:1.15 as builder

# RUN mkdir -p /myapp
# ADD . /myapp
# WORKDIR /myapp

# # Собираем бинарный файл
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
#    go build -o /myapp .

# FROM scratch

# COPY --from=builder /myapp /myapp

# CMD ["/myapp"]




# # #4 Cert
# FROM golang:1.15 as builder

# RUN mkdir -p /myapp
# ADD . /myapp
# WORKDIR /myapp

# # Собираем бинарный файл
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
#    go build -o /myapp .

# FROM scratch

# COPY --from=builder /myapp /myapp
# COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

# CMD ["/myapp"]




# #5 Non-root user
# FROM golang:1.15 as builder

# RUN mkdir -p /myapp
# ADD . /myapp
# WORKDIR /myapp

# # Добавляем непривилегированного пользователя
# RUN useradd -u 10001 myapp

# # Собираем бинарный файл
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
#    go build -o /myapp ./cmd/myapp

# FROM scratch

# # Не забываем скопировать /etc/passwd с предыдущего стейджа
# COPY --from=builder /etc/passwd /etc/passwd
# USER myapp

# COPY --from=builder /myapp /myapp
# COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

# CMD ["/myapp"]






#6 With modules
FROM golang:1.15 as modules

ADD go.mod go.sum /m/
RUN cd /m && go mod download

FROM golang:1.15 as builder

COPY --from=modules /go/pkg /go/pkg

RUN mkdir -p /myapp
ADD . /myapp/
WORKDIR /myapp

# Добавляем непривилегированного пользователя
RUN useradd -u 10001 myapp

# Собираем бинарный файл
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
   go build -o /myapp ./cmd/myapp

FROM scratch

# Не забываем скопировать /etc/passwd с предыдущего стейджа
COPY --from=builder /etc/passwd /etc/passwd
USER myapp

COPY --from=builder /myapp /myapp
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/

EXPOSE 8080 

CMD ["/myapp"]