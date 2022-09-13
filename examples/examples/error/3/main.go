package main

import (
	"errors"
)

type NotFoundError struct {
    Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }


type QueryError struct {
    Query string
    Err   error
}

func (e *QueryError) Unwrap() error { return e.Err }

//---------------------
В Go 1.13 функция fmt.Errorf поддерживает новая команда %w. Если она есть, то ошибка, возвращаемая fmt.Errorf, будет содержать метод Unwrap, возвращающий аргумент %w, который должен быть ошибкой. Во всех остальных случаях %w идентична %v.

if err != nil {
    // Return an error which unwraps to err.
    return fmt.Errorf("decompress %v: %w", name, err)
}


Упаковка ошибки с помощью %w делает её доступной для errors.Is и errors.As:

err := fmt.Errorf("access denied: %w", ErrPermission)
...
if errors.Is(err, ErrPermission)
//---------------------

//Функция errors.Is сравнивает ошибку со значением.

// Similar to:
//   if err == ErrNotFound { … }
if errors.Is(err, ErrNotFound) {
    // something wasn't found
}



//Функция As проверяет, относится ли ошибка к конкретному типу.

// Similar to:
//   if e, ok := err.(*QueryError); ok { … }
var e *QueryError
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}