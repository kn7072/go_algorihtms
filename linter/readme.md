## Установка

https://golangci-lint.run/usage/install/
```bash
In alpine linux (as it does not come with curl by default)
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.49.0

golangci-lint --version
```
### Добавляем строки в файл ~/.profile 
```bash
# GO LINTER
export PATH=/home/stapan/go/bin:$PATH
```

### Настройка visual studio
- нажимаем f1
- печатаем open settings
- выбираем из меню open settings (JSON)

дописываем 
"go.lintTool":"golangci-lint",
    "go.lintFlags": [
        "--fast"
    ]
ИЛИ указываем путь к .golangci.yml чтобы использовать для всех проектов
"go.lintFlags": [
        "--fast",
        "-c",
        "~/go/LINTER/.golangci.yml"
    ]
~/go/LINTER/.golangci.yml путь до файла конфигурации

### добавляем файл конфигурации .golangci.yml

### Полезные ссылки
- https://habr.com/ru/post/457970/
- https://golangci-lint.run/usage/integrations/

### Команды
golangci-lint --version
golangci-lint run --disable-all -E revive -E errcheck -E nilerr -E gosec