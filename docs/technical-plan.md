# Technical Plan for AltVerse

Этот документ описывает структуру папок и ключевые решения, которые были перечислены в конспекте. Основные каталоги для бэкэнда включают `gameplay-engine`, `api-gateway`, `services/analytics`, `services/anti-cheat` и `services/payments`.

Каждый сервис содержит минимальный набор файлов (`main`, конфигурации, Dockerfile). Подробную структуру см. ниже.

```
backend/
├─ gameplay-engine/
│  ├─ cmd/engine/main.go
│  ├─ internal/{game,network,store,util}
│  ├─ proto/gameplay.proto
│  └─ Makefile
├─ api-gateway/
│  ├─ src/{main.ts,modules,...}
│  ├─ test/e2e
│  └─ package.json
├─ services/
│  ├─ analytics/
│  ├─ anti-cheat/
│  └─ payments/
```

Эта структура отображает разделение ответственности и будет расширяться по мере разработки. Подробности см. в основном описании концепции.
