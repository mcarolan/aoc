package util

import "golang.org/x/exp/constraints"

func Map[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))

	for i, t := range s {
		result[i] = f(t)
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
