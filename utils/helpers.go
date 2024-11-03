package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// get a path param from request
func GetPathParam(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	value, ok := vars[name]
	if !ok {
		return "", fmt.Errorf("invalid or missing %s in request param", name)
	}
	return value, nil
}

// MapConcurrent performs concurrent mapping using goroutines
func MapConcurrent[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, len(items))
	ch := make(chan struct{}, len(items)) // Buffered channel for concurrent operations

	for i, item := range items {
		go func(index int, val T) {
			result[index] = fn(val)
			ch <- struct{}{} // Signal completion
		}(i, item)
	}

	// Wait for all goroutines to complete
	for i := 0; i < len(items); i++ {
		<-ch
	}

	return result
}
