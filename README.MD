# METROID BOOKMARKS

<img alt="text" height="100" src="repository_files/metroid-svgrepo-com.svg" width="100"/>

## DEVELOP

### 1. Запуск

1. Указать тестовые параметры БД и прочего в [config.json](backend/etc/config/config.json). Параметр "debug"
   оставить в значении true.
2. Создать контейнеры бекенда:

```bash
cd backend
docker compose up
```

Во время разработки используется утилита air в docker сервисе "web".  
Air автоматически подгрузит библиотеки из go.mod в контейнер, даже если добавить новые во время разработки.

### 2. Обновление swagger

Если внести изменения в синтаксис библиотеки swaggo(комментарии к [эндпоинтам](backend/pkg/handler/auth.go#L12) или
к [функции main](backend/cmd/main.go#L19)), то необходимо
обновить swagger.

1. Установить swagger:

```bash
cd backend
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Создать/обновить swagger:

```bash
cd backend
swag init -g cmd/main.go
```

_Если использовать Docker, то команды выше исполняются так:_

```bash
cd backend
sudo docker compose exec web go install github.com/swaggo/swag/cmd/swag@latest
sudo docker compose exec web swag init -g cmd/main.go
```

3. Выравнивание структур

```bash
fieldalignment --fix ./...
```

## PRODUCTION

### 1. Обновление Backend

1. Сменить тестовые параметры на реальные в [config.json](backend/etc/config/config.json). Параметр "debug" равен false.

2. Удалить проксирование портов службы web в [docker-compose.yml](backend/docker-compose.yml#L7.

3. Выполнить команду обновления контейнеров:

 ```bash
 docker compose up
 ```

### 2. Обновление Frontend

1. Скопировать директорию [deploy](server_files/deploy) в репозиторий фронтенд приложения:

```bash
cp -r server_files/deploy/ /frontend/deploy
```

2. Создать и запустить контейнер фронтенда:

```bash
cd frontend
sudo docker build -t front -f ./.deploy/docker/Dockerfile .
sudo docker run -p 80:80 --name front --network backend_default front
```
