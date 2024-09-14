# METROID BOOKMARKS

<img alt="text" height="100" src="repository_files/metroid-svgrepo-com.svg" width="100"/>

### BACKEND

Документация к [Backend сервису](./backend/README.MD)

## NGINX и FRONTEND

### 1. Обновление Frontend

1. Скопировать директорию [deploy](./repository_files/deploy) в репозиторий фронтенд приложения:

```bash
cp -r repository_files/deploy/ /frontend/deploy
```

2. Создать и запустить контейнер фронтенда:

```bash
cd frontend
sudo docker build -t front -f ./.deploy/docker/Dockerfile .
sudo docker run -p 80:80 --name front --network backend_default front
```
