## Основная цель

Напишите сервис юного астролога.
Сервис должен выполнять основные функции:
1) При помощи воркера раз в сутки получать метаданные и картинку дня [APOD](https://api.nasa.gov/), сохранять в реализованное любым способом бинарное хранилище;
2) Реализовывать Http API сервер: метод получения всех записей альбома и записи за выбранный день;
3) Сервис должен собираться в docker-образ;

**Технологии**: postgreSQL, docker, docker-compose, makefile

**Примечания:**
- Для конфигурирования проекта используйте переменные ENV;
- Сервис должен создать базу данных и необходимую структуру таблиц при запуске;