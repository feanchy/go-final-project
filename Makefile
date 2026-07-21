APP_NAME=scheduler
PORT=7540


.PHONY: build test run docker-build docker-run clean


# Сборка Go приложения
build:
	go build -o $(APP_NAME) .


# Запуск тестов
test:
	go test ./tests


# Локальный запуск
run:
	go run .


# Сборка Docker образа
docker-build:
	docker build -t $(APP_NAME) .


# Запуск Docker контейнера
docker-run:
	docker run --rm -p $(PORT):$(PORT) $(APP_NAME)


# Удалить бинарник
clean:
	rm -f $(APP_NAME)