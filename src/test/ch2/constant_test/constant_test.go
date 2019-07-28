package constant_test

import "testing"

const (
	Monday = iota +1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
const (
	Open = 1 << iota
	Close
)
func TestConstantTry(t *testing.T)  {
	t.Log(Monday, Tuesday)
}
func TestConstantTryBin(t *testing.T)  {
	a := 7 //0111
	t.Log( a&Open == Open, a&Close == Close)
}