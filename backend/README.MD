# Metroid Bookmarks Backend

У проекта есть режимы Develop(разработка) и Production(рабочий режим).
В режиме Develop запускается вотчер, который перекомпилирует проект при изменении контрольной суммы любого go файла.
На стартапе вотчер запускает задачи заданные в [конфигурации](/conf/.air.toml).
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

3. Запустите приложение командой:

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
Генерация свагер документации включена в вотчер.

## PRODUCTION

1. В .env переменную окружения установить в значение PRODUCTION=true.
2. Соберите проект командой:

```bash
    make build
```

3. Запустите приложение командой:

```bash
    make start
```

P.s. Для других команд смотрите [Makefile](Makefile).