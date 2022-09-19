package main

// https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#multiple-type-parameters

type Stringer interface {
	String() string
}

type Plusser interface {
	Plus(string) string
}

func ConcatTo[S Stringer, P Plusser] (s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}

	return r
}

// A single constraint can be used for multiple type parameters, 
// just as a single type can be used for multiple non-type function parameters. 
// The constraint applies to each type parameter separately.

// Stringify2 converts two slices of different types to strings,
// and returns the concatenation of all the strings.
func Stringify2[T1, T2 Stringer](s1 []T1, s2 []T2) string {
	r := ""

	for _, v1 := range s1 {
		r += v1.String()
	}

	for _, v2 := range s2 {
		r += v2.String()
	}
		return r
}