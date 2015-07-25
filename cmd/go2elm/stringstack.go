package main

type StringStack []string

func (self *StringStack) push(s string) {
	*self = append(*self, s)
}

func (self *StringStack) pop() {
	*self = (*self)[0 : len(*self)-1]
}

func (self *StringStack) with(s string, fn func()) {
	self.push(s)
	fn()
	self.pop()
}

func (self StringStack) top() string {
	return self[len(self)-1]
}
