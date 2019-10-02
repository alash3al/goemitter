package Emitter

import (
	"reflect"
	"testing"
)

func TestRemoveListener(t *testing.T) {
	emitter := Construct()

	counter := 0
	fn1 := func(args ...interface{}) {
		counter++
	}
	fn2 := func(args ...interface{}) {
		counter++
	}

	emitter.On("testevent", fn1)
	emitter.On("testevent", fn2)

	emitter.RemoveListener("testevent", fn1)
	emitter.EmitSync("testevent")

	listenersCount := emitter.ListenersCount("testevent")

	expect(t, 1, listenersCount)
	expect(t, 1, counter)
}

func TestOnce(t *testing.T) {
	emitter := Construct()

	counter := 0
	fn := func(args ...interface{}) {
		counter++
	}

	emitter.Once("testevent", fn)

	emitter.EmitSync("testevent")
	emitter.EmitSync("testevent")

	expect(t, 1, counter)
}

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
