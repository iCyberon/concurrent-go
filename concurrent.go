package concurrent

import (
	"fmt"
	"reflect"
	"sync"
)

// concurrentGroup
type concurrentGroup struct {
	functions []reflect.Value
	arguments [][]reflect.Value
	errors    []error
}

// New .
// TODO: Add decription
func New() *concurrentGroup {
	cg := &concurrentGroup{}
	return cg
}

// Add .
// TODO: Add decription
func (c *concurrentGroup) Add(fn interface{}, args ...interface{}) *concurrentGroup {
	method := reflect.ValueOf(fn)
	arguments := make([]reflect.Value, len(args))

	if method.Kind() != reflect.Func {
		c.errors = append(c.errors, fmt.Errorf("%v is not a function", fn))
		return c
	}

	for i, arg := range args {
		arguments[i] = reflect.ValueOf(arg)
	}

	c.functions = append(c.functions, method)
	c.arguments = append(c.arguments, arguments)

	return c
}

// Exec .
// TODO: Add decription
func (c *concurrentGroup) Exec() ([][]interface{}, []error) {
	if len(c.errors) > 0 {
		return nil, c.errors
	}

	wg := &sync.WaitGroup{}
	results := make([][]interface{}, len(c.functions))

	for i, m := range c.functions {
		args := c.arguments[i]

		wg.Add(1)
		go func(method reflect.Value, args []reflect.Value, index int) {
			res := method.Call(args)
			values := make([]interface{}, len(res))

			for ix, val := range res {
				values[ix] = val.Interface()
			}

			results[index] = values
			wg.Done()
		}(m, args, i)
	}

	wg.Wait()

	return results, nil
}
