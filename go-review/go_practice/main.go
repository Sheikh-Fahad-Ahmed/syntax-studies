package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"math"
	"slices"
	"time"
	"unicode/utf8"
)

func main() {
	// fmt.Println("Hello, World!")

	// Values  ----------------------------------
	// values()

	// Variables  ----------------------------------
	// variables()

	// Constants  ----------------------------------
	// constants()

	// Looping  ----------------------------------
	// looping()

	// if-Else  ----------------------------------
	// ifElse()

	// Switch  ----------------------------------
	// switchCase()

	// arrays  ----------------------------------
	// arrays()

	// Slices  ----------------------------------
	// sliceInfo()

	// Maps  ----------------------------------
	// mapping()

	// Functions  ----------------------------------
	// res := plus(1, 2)
	// fmt.Println(res)

	// res = plusPlus(3, 4, 5)
	// fmt.Println(res)

	// Multiple Return types ----------------------------------
	// a, b := vals()
	// fmt.Println(a, b)

	// _, c := vals()
	// fmt.Println(c)

	// Variadic Functions ----------------------------------
	// variadicFuncSums(1, 2)
	// variadicFuncSums(1, 2, 3)

	// nums := []int{1, 2, 3, 4}
	// variadicFuncSums(nums...)

	// Closures  ----------------------------------
	// intFunc := closures()

	// fmt.Println(intFunc())
	// fmt.Println(intFunc())

	// Recursion ----------------------------------
	// fmt.Println(fact(7))

	// var fib func(n int) int

	// fib = func(n int) int {
	// 	if n < 2 {
	// 		return n
	// 	}
	// 	return fib(n-1) + fib(n-2)
	// }

	// fmt.Println(fib(7))

	// Range over Built-in Types ----------------------------------
	// rangeOverBuiltInTypes()

	// Pointers ----------------------------------
	// pointers()

	// String and Runes
	// stringsAndRunes()

	// Structs
	// structs()

	// JSON
	jsonFunc()

}

func values() {

	fmt.Println("go" + "lang")
	fmt.Println("1+1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func variables() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero value
	fmt.Println(e)

	f := "banana"
	fmt.Println(f)

}

const s string = "constant"

func constants() {
	fmt.Println(s)

	const n = 50000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}

func looping() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j <= 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range: ", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}

func ifElse() {
	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 10%2 == 0 || 5%2 == 0 {
		fmt.Println("either 10 or 5 is even")
	}

	if num := 9; num < 0 {
		fmt.Println("num is negative")
	} else if num < 10 {
		fmt.Println("num has 1 digit")
	} else {
		fmt.Println("num has multiple digits")
	}
}

func switchCase() {
	i := 2
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("its the weekday, back to work!")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("Its afternoon")
	}

	whatAmiI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("im a boolean")
		case int:
			fmt.Println("Im an integer")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}
	whatAmiI(true)
	whatAmiI(10)
	whatAmiI("hi")

}

func arrays() {
	// Fixed size, slices are more common
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println("dcl", b)

	b = [...]int{100, 3: 400, 600}
	fmt.Println("idx:", b)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{3, 4, 5},
	}
	fmt.Println("2d:", twoD)

}

func sliceInfo() {
	var s []string
	fmt.Println("unini:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get", s[1])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "z")
	fmt.Println("apd:", s)

	c := make([]string, 3)
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d", twoD)
}

func mapping() {
	m := make(map[string]int)

	m["a1"] = 1
	m["a2"] = 2

	fmt.Println("map:", m)

	v1 := m["a1"]
	fmt.Println("v1:", v1)

	v3 := m["a3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "a2")
	fmt.Println("map", m)

	clear(m)
	fmt.Println("cleared map:", m)

	_, ok := m["a2"]
	if !ok {
		fmt.Println("a2 is empty")
	}

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func variadicFuncSums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func closures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func rangeOverBuiltInTypes() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func pointers() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("Pointer:", &i)
}

func stringsAndRunes() {
	a := "hello"

	fmt.Println("len:", len(a))

	const s = "สวัสดี"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune Count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Printf("\nUsing DecodeRuneInString\n")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	return &person{
		name: name,
		age:  42,
	}
}

func structs() {
	fmt.Println(person{"bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("John"))

	s := person{"Sean", 21}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp)

	sp.age = 51
	fmt.Println(s.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func jsonFunc() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(intB)

	fltB, _ := json.Marshal(3.12)
	fmt.Println(fltB)

	strB, _ := json.Marshal("gopher")
	fmt.Println(strB)

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 3}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]any

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[1])

	
}
