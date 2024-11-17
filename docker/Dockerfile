FROM ubuntu:24.04
RUN mkdir /app
WORKDIR /app
COPY  ./cmd  .
#ENTRYPOINT [ "/app","defaultvalue" ]
# FROM golang:1.20 AS builder
# FROM ubuntu:24.04
# WORKDIR /app
# RUN mkdir /app
# COPY . .
# RUN go build -o main ./cmd/main.go
# COPY  ./cmd  .

# FROM ubuntu:22.04

# # Устанавливаем рабочую директорию
# WORKDIR /app

# # Копируем бинарный файл из этапа сборки
# COPY --from=builder /app/main .

# # Копируем директорию хранения данных (пустую, данные будут загружаться в runtime)
# COPY --from=builder /app/internal/pkg/storage/data ./internal/pkg/storage/data


# # Устанавливаем команду запуска
# CMD ["./main"]