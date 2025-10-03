Сервис для управления дискордом и для дикорд бота на ElysiumSMP

---

Сборка в Docker
```bash
docker build . -t svc-discord:latest
docker run --rm svc-discord:latest
```

Для конфигурации используются вот такие ENV переменные:
- DISCORD_BOT_TOKEN - токен бота в дискорд
