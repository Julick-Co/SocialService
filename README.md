# SocialService
Сервис для отношений

## Генерация protobuf и gRPC

Запусти команды в терминале:

```bash
# Установка генераторов Go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

```bash
# Генерация кода из proto-файлов
protoc --proto_path=./ --go_out=. --go-grpc_out=. api/social/v1/messages.proto api/social/v1/service.proto
