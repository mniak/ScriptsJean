package main

func ptr[T any](value T) *T {
	return &value
}
