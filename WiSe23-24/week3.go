// Woche 3

package main

import "fmt"

func hello3(x int) func() {
	return func() { fmt.Printf("Hello %d \n", x) }
}

func testHello() {

	// Verschiedenen Anwendungsformen

	// 1. Partielle Funktionsanwendung

	var f func()

	f = hello3(3)

	f()
	// 2.
	h := hello3(3)

	h()

	// 3.
	(hello3(2))()
	//              ^^^^^^^^^^
	//              ist eine Funktion welche direkt aufgerufen wird

	// 4. Funktionsanwendung ist links-assoziativ

	hello3(2)()
}

// Addieren

func add(x int) func(int) int {
	return func(y int) int { return x + y }
}

func plus(x int, y int) int {
	return x + y
}

func testAdd() {
	fmt.Printf("%d", plus(3, 4))

	fmt.Printf("%d", add(3)(4))

	inc := add(1)

	fmt.Printf("%d", inc(4))

	g := add
	fmt.Printf("%d", g(3)(2))

	var f func(int) func(int) int
	f = add
	fmt.Printf("%d", f(3)(2))

}

func main() {

	// testHello()
	testAdd()
}

/*

Frage aus der Vorlesung.

Wäre folgende Syntax gültig: curry(f func(int, int) int) func(int)(int) int


Nein!

func(int)(int) int

ist syntaktisch nicht legal.

es fehlt "func"

 func(int)( func(int) int )

*/
