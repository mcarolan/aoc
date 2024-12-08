package util

import "golang.org/x/exp/constraints"

func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))

	for i, t := range s {
		result[i] = f(t)
	}

	return result
}

func FlatMap[T, U any](s []T, f func(T) []U) []U {
	return Flatten(Map(s, f))
}

func Flatten[T any](s [][]T) []T {
	result := make([]T, 0)

	for _, group := range s {
		result = append(result, group...)
	}

	return result
}

func Filter[T any](s []T, p func(T) bool) []T {
	result := make([]T, 0)

	for _, t := range s {
		if p(t) {
			result = append(result, t)
		}
	}

	return result
}

func ForAll[T any](s []T, p func(T) bool) bool {
	result := true

	for _, t := range s {
		if !p(t) {
			return false
		}
	}

	return result
}

func Sum[T constraints.Integer](s []T) T {
	var result T
	for _, t := range s {
		result += t
	}
	return result
}

func Pairs[T any](s []T) <-chan [2]T {
	ch := make(chan [2]T)

	go func() {
		for i := 0; i < len(s); i++ {
			for j := i + 1; j < len(s); j++ {
				ch <- [2]T{s[i], s[j]}
			}
		}
		close(ch)
	}()

	return ch
}
