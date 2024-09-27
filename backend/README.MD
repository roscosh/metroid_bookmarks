# Metroid Bookmarks Backend

У проекта есть режимы Develop(разработка) и Production(рабочий режим). 
В режиме Production закрыты порты у контейнеров postgreSQL, redis, app. Доступ к ним возможен только networks, у каждого 
режима он задан в конце его docker-compose.yml. В режиме Develop все порты открыты. 

Также отличием режима Develop от Production является запуск вотчера, который перекомпилирует проект и выполнит остальные
задачи, описанные в [конфигурации утилиты Air](/conf/.air.toml), при изменении контрольной суммы любого go файла.

Остальные различия смотрите в [конфигирации контейнеров](/conf/docker).

Для запуска проекта нужно:

1. создать файл .env в корне проекта, и заполнить его переменными из примера(.env.example).
2. Заполнить в [конфиге приложения](/conf/app_config.json) чаты для отправки.

## DEVELOP

### Запуск

1. В .env переменную окружения установить в значение PRODUCTION=false.
2. Соберите проект командой:

```bash
    make build
```

3. Выполните команду миграции БД:

```bash
    make migrate
```

4. Запустите приложение командой:

```bash
    make start
```

### Инструменты для разработки

1. Выравнивание структур

```bash
go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@latest
fieldalignment ./...
```

2. Linter

```bash
   make lint
```

3. Тестирование

```bash
   make tests
```

### Endpoints

Все эндпоинты описаны в swagger. Адрес swagger: http://localhost:3000/swagger/index.html.
Генерация свагер документации включена в вотчер. Не доступно в режиме PRODUCTION.

## PRODUCTION

### Запуск

1. В .env переменную окружения установить в значение PRODUCTION=true.
2. Соберите проект командой:

```bash
    make build
```

3. Выполните команду миграции БД:

```bash
    make migrate
```

4. Запустите приложение командой:

```bash
    make start
```

P.s. Для других команд смотрите [Makefile](Makefile).