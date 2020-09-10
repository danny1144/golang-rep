package main

import (
	"fmt"
	"reflect"
	"strconv"
	"sync"
)

func main() {
	//Go各种类型转换及函数的高级用法
	//字符串转整形
	var changeStr = "1234"
	fmt.Println(strconv.ParseInt(changeStr, 10, 0))

	//golang不支持隐式类型转换
	var tint byte = 100
	fmt.Println(int(tint))
	//int转字符串
	convStr := strconv.Itoa(4)
	fmt.Println(convStr, reflect.TypeOf(convStr))

	//window int默认类型
	var i int = 43
	var float64 = float64(i)
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(float64))

	// 数组初始化
	arr1 := [3]int{2, 2, 3}
	arr2 := [5]int{0: 3, 3: 4, 4: 6}
	arr3 := [...]string{"sdf", "aaa", "sdf sdf "}
	arr4 := [...]struct {
		name string
		age  uint
	}{
		{"张三", 35},
		{"李四", 32},
		{"王五", 14},
	}

	fmt.Println(arr1, arr2, arr3, arr4)
	//rune int32的别名，Unicode码用utf-8进行编码
	str := len(string(rune('好')))
	fmt.Println(str)
	//获取中文字符串的长度，需要转换成rune切片
	strZh := "中国人"
	fmt.Println(len([]rune(strZh)))

	go fmt.Println("开启邪趁")
	//数组是值传递
	x := [3]int{1, 2, 3}
	func(arr [3]int) {
		arr[0] = 5
		fmt.Println(arr)
	}(x)
	fmt.Println(x)
	//map遍历是顺序不固定的
	m := map[string]string{
		"1": "3",
		"2": "4",
		"3": "5",
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
	fmt.Println(f(3, 4))
	fmt.Println(Sum(2, 3, 4, 5, 6, 4, 3))
	fmt.Println(Find(map[int]int{2: 2, 3: 5}, 2))
	//接口
	var a = []interface{}{123, "abc"}
	fmt.Println(a...)
	fmt.Println(Inc(3))

	for i := 0; i < 3; i++ {
		i := i
		//Possible resource leak, 'defer' is called in a 'for' loop
		defer func() { println(i) }()
	}
	//通过函数传入i，调用defer语句马上执行求值
	for i := 0; i < 3; i++ {
		defer func(i int) { println(i) }(i)
	}
	println(">>>>>>>>>>>>>>>>>>>>>>>")
	var arr = []int{2, 3, 5, 3, 3}
	Twice(arr)
	for i, v := range arr {
		println(i, v)
	}

	var wg sync.WaitGroup

	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)

}

//用for range方式迭代的性能可能会更好一些，因为这种迭代可以保证不会出现数组越界的情形，每轮迭代对数组元素的访问时可以省去对下标越界的判断。
//匿名函数
var f = func(a, b int) int {
	return a + b
}

//Sum is a function
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	return a
}

// Find is search elements in map
func Find(m map[int]int, key int) (value int, ok bool) {
	value, ok = m[key]
	return
}

// Inc test defer
func Inc(v int) int {

	defer func() { v++ }()
	return v
}

// Twice make erver elements double
func Twice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

//IntSliceHeader int切片结构体
type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

//TwiceStruct double element
func TwiceStruct(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}
