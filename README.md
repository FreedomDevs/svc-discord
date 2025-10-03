Сервис для управления дискордом и для дикорд бота на ElysiumSMP

---

Сборка в Docker
```bash
docker build . -t svc-discord:latest
docker run --rm svc-discord:latest
```

Для конфигурации используются вот такие ENV переменные:
- DISCORD_BOT_TOKEN - токен бота в дискорд
- WARN_LEVELS - через запятую ID ролей для варнов
- WARN_ACCESS_ROLES - через запятую ID ролей которым доступен варн
- GUILD_ID - айди дискорд сервера на котором бот будет работать
