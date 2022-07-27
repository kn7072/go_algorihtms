// https://habr.com/ru/company/karuna/blog/552944/

package set

type Set[T comparable] map[T]struct{}

func Make[T comparable]() Set[T] {
    return make(Set[T])
}

func (s Set[T]) Add(v T) {
    s[v] = struct{}{}
}

func (s Set[T]) Delete(v T) {
    delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
    _, ok := s[v]
    return ok
}

func (s Set[T]) Len() int {
    return len(s)
}

func (s Set[T]) Iterate(f func(T)) {
    for v := range s {
        f(v)
    }
}