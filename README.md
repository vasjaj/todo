# Application
Переделал прошлое тестовое задание, оставил makefile, docker и прочие уже настроенные вещи.
## Project strutrure
Project structure is based on https://github.com/golang-standards/project-layout
- `/build` - Dockerfile
- `/cmd` - main.go
- `/configs` - примем конфига, sql скрипт
- `/internal` - internal packages
  - `/config` - configuration package
  - `/database` - database package
  - `/server` - server package
## Testing
Added test example in `internal/server/seamless_test.go`
## Comments
  - При запросе withdrawAndDeposit создатеся запись в таблице transactions. Создание тразнакции происходит в транзакции, чтобы не было проблем с параллельными запросами.
  - При запросе getBalance баланас игрока считается на основании прошлых транзакций в конкретной валюте. Это сделано чтобы не хранить текущий баланс в базе и не нарушать нормализацию данных.
