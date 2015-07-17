package possessed

import (
	"reflect"
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

/// POSSESSS

func Test_Possess_EmptyString(t *testing.T) {
	s := Possess("")
	expect(t, s, "")
}

func Test_Possess_WhitespaceyString(t *testing.T) {
	s := Possess("  ")
	expect(t, s, "")

	s = Possess("  \r\n")
	expect(t, s, "")
}

func Test_Possess_Gladys(t *testing.T) {
	s := Possess("Gladys")
	expect(t, s, "Gladys'")
}

func Test_Possess_Dave(t *testing.T) {
	s := Possess("Dave")
	expect(t, s, "Dave's")
}

func Test_Possess_It(t *testing.T) {
	s := Possess("it")
	expect(t, s, "its")

	s = Possess("It")
	expect(t, s, "Its")
}

/// UNPOSSESSS

func Test_Unpossess_EmptyString(t *testing.T) {
	s := Unpossess("")
	expect(t, s, "")
}

func Test_Unpossess_WhitespaceyString(t *testing.T) {
	s := Unpossess("  ")
	expect(t, s, "")

	s = Unpossess("  \r\n")
	expect(t, s, "")
}

func Test_Unpossess_Gladys(t *testing.T) {
	s := Unpossess("Gladys'")
	expect(t, s, "Gladys")

	s = Unpossess("Gladys’")
	expect(t, s, "Gladys")

	s = Unpossess("Gladys")
	expect(t, s, "Gladys")
}

func Test_Unpossess_Dave(t *testing.T) {
	s := Unpossess("Dave's")
	expect(t, s, "Dave")

	s = Unpossess("Dave’s")
	expect(t, s, "Dave")

	s = Unpossess("Daves")
	expect(t, s, "Daves")
}

func Test_Unpossess_It(t *testing.T) {
	s := Unpossess("its")
	expect(t, s, "it")

	s = Unpossess("Its")
	expect(t, s, "It")
}
