go install golang.org/x/tools/cmd/stringer


//go:generate ./command.sh

//go:generate -command ls -l
//go:generate -command bye echo "Goodbye, world!"  -создаем псевдоним команды bye

//go:generate bye -выполняем команду по псевдониму
//go:generate go run generate.go

go generate -v показывает какие файлы будут обрабатываться
go generate -x показывает команды которые должны выполняться и выполняет их

go generate -n посмотреть команды, но не выполнять их

go generate -run bye выполнить команду по всевдониму

