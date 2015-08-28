// An EventEmitter package for go 1.2 +
// Author: Mohammed Al Ashaal
// License: MIT License
// Version: v1.0.0
package Emitter

// Emitter - our listeners container
type Emitter struct {
	listeners	map[interface{}][]Listener
}

// Listener - our callback container and whether it will run once or not
type Listener struct {
	callback	func(...interface{})
	once		bool
}

// Construct() - create a new instance of Emitter
func Construct() *Emitter {
	emitter := new(Emitter)
	emitter.listeners = make(map[interface{}][]Listener)
	return emitter
}

// Destruct() - free memory from an emitter instance
func(self *Emitter) Destruct() {
	self = nil
} 

// AddListener() - register a new listener on the specified event
func(self *Emitter) AddListener(event string, callback func(...interface{})) *Emitter {
	return self.On(event, callback)
}

// On() - register a new listener on the specified event
func(self *Emitter) On(event string, callback func(...interface{})) *Emitter {
	if _, ok := self.listeners[event]; ! ok {
		self.listeners[event] = []Listener{}
	}
	self.listeners[event] = append(self.listeners[event], Listener{callback, false})
	self.EmitSync("newListener", []interface{}{event, callback})
	return self
}

// Once() - register a new one-time listener on the specified event
func(self *Emitter) Once(event string, callback func(...interface{})) *Emitter {
	if _, ok := self.listeners[event]; ! ok {
		self.listeners[event] = []Listener{}
	}
	self.listeners[event] = append(self.listeners[event], Listener{callback, true})
	self.EmitSync("newListener", []interface{}{event, callback})
	return self
}

// RemoveListeners() - remove the specified callback from the specified events' listeners
func(self *Emitter) RemoveListener(event string, callback func(...interface{})) *Emitter {
	if _, ok := self.listeners[event]; ! ok {
		return self
	}
	for k,v := range(self.listeners[event]) {
		if &v.callback == &callback {
			self.listeners[event] = append(self.listeners[event][:k], self.listeners[event][k + 1:]...)
			self.EmitSync("removeListener", []interface{}{event, callback})
			return self
		}
	}
	return self
}

// RemoveAllListeners() - remove all listeners from (all/event)
func(self *Emitter) RemoveAllListeners(event interface{}) *Emitter {
	if event == nil {
		self.listeners = make(map[interface{}][]Listener)
		return self
	}
	if _, ok := self.listeners[event]; ! ok {
		return self
	}
	delete(self.listeners, event)
	return self
}

// Listeners() - return an array with the registered listeners in the specified event
func(self *Emitter) Listeners(event string) []Listener {
	if _, ok := self.listeners[event]; ! ok {
		return nil
	}
	return self.listeners[event]
}

// ListenersCount() - return the count of listeners in the speicifed event
func(self *Emitter) ListenersCount(event string) int {
	return len(self.Listeners(event))
}

// EmitSync() - run all listeners of the specified event in synchronous mode
func(self *Emitter) EmitSync(event string, args ...interface{}) *Emitter {
	for _,v := range(self.Listeners(event)) {
		if v.once {
			self.RemoveListener(event, v.callback)
		}
		v.callback(args...)
	}
	return self
}

// EmitAsync() - run all listeners of the specified event in asynchronous mode using goroutines
func(self *Emitter) EmitAsync(event string, args []interface{}) *Emitter {
	for _,v := range(self.Listeners(event)) {
		if v.once {
			self.RemoveListener(event, v.callback)
		}
		go v.callback(args...)
	}
	return self
}
