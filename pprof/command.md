- Как видно, мы просто передаем утилите pprof путь к хендлеру, по которому «слушает» профайлер. 
    Дополнительно можно передать время работы профайлера (по умолчанию 30 секунд).

    $ go tool pprof http://localhost:8080/debug/pprof/profile?seconds=5

- С помощью команды list можно подробно исследовать каждую функцию, например, list leftpad:
    (pprof) list leftpad

- Профилирование кучи
    go tool pprof goprofex http://127.0.0.1:8080/debug/pprof/heap
    По умолчанию он показывает объём используемой памяти:

    Но нас больше интересует количество размещённых в куче объектов. Запустим pprof с опцией -alloc_objects:
    go tool pprof -alloc_objects goprofex http://127.0.0.1:8080/debug/pprof/heap

    анализируем сгенерированный файл
    go tool pprof pprof.goprofex.alloc_objects.alloc_space.inuse_objects.inuse_space.002.pb.gz
    (pprof) top

- выкачаем то, что возвращает URL самостоятельно:

    $ curl http://localhost:8080/debug/pprof/profile?seconds=5 -o /tmp/cpu.log

    Мы видим, что внутри /tmp/cpu.log такие же бинарные данные, какие возвращаются при использовании go tool test -cpuprofile или StartCPUProfile(). «Натравим» команду strings на этот бинарный файл и поймем, что внутри нет названий функций или так называемых символов.

    $ strings /tmp/cpu.log | grep cpuhogger

    Откуда же тогда в первом случае, когда мы запускали pprof без бинарника, были получены имена функций? Оказывается, при импорте net/http/pprof добавляется еще один URL /debug/pprof/symbol, который по адресу функции возвращает ее название. С помощью запросов к этому URL команда pprof получает имена функций.

- запускаем сбор данных профайлинга CPU, указав кол-во секунд, в течение которых мы профилируем:
    curl http://localhost:8080/debug/pprof/profile?seconds=5 > ./profile

    go tool pprof ./profile 

    (pprof) top15 - посмотрим топ 15 функций по времени выполнения:

    Мы можем посмотреть исходный код функции и в нем будут показано время на отдельные вызовы.
    (pprof) list text/template.\(\*state\).walk$

- (pprof) pdf -сохраняется граф в pdf 
- Команда top10 --cum возвращает совокупное время выполнения для каждой функции:
    (pprof) top10 --cum

- (pprof) svg -сохраняется граф в svg

-  профилирования с использованием веб-интерфейса
go tool pprof -http=localhost:8090 ./profile



https://tproger.ru/translations/memory-leaks-investigation-in-go-using-pprof/
Type: inuse_space
Time: Jan 22, 2019 at 1:08pm (IST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)

Здесь важно отметить Type: inuse_space. Мы смотрим на данные выделения памяти в определённый момент (когда мы захватили профиль). Тип является значением конфигурации sample_index, а возможными значениями могут быть:

    inuse_space — объём выделенной и ещё не освобождённой памяти;
    inuse_objects — количество выделенных и ещё не освобождённых объектов;
    alloc_space — общий объём выделенной памяти (независимо от освобождённой);
    alloc_objects — общее количество выделенных объектов (независимо от освобождённых).

можно переключаться между режимами
(pprof) alloc_space
(pprof) function_name

Теперь введите top в интерактивной консоли, и на выводе будут главные потребители памяти.

В выводимом списке можно увидеть два значения — flat и cum.

    flat означает, что память выделена функцией и удерживается ей;
    cum означает, что память выделена функцией или функцией, которая была вызвана стеком.

https://habr.com/ru/company/badoo/blog/301990/
