package main
import "fmt"
func main() {
	/**
	基本数据类型
	 */
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("---------------------------------------------------")
	/**
	变量
	*/
	var a string = "土拨鼠" //var 声明变量，也可以申请多个
	fmt.Println(a)

	var b, c int = 1,2  //一次性声明多个变量。
	fmt.Println(b,c)

	var d = true        // Go 将自动推断已经初始化的变量类型。
	fmt.Println(d)

	e := "blockchain"  // := 语句是申明并初始化变量的简写，等价于 var f string = "blockchain"。
	fmt.Println(e)

	/**
	  基本类型的默认值
	  int --------> 0
	  string --------> " "
	  bool --------> false
	  float --------> 0.0
	*/
	var f int          // 声明变量且没有给出对应的初始值时，变量将会初始化为默认值。
	fmt.Println(f)

	g, _ := 100, 200 // _标识匿名变量，这里第二个值200赋给了_， 表示后续代码不需要再用此变量
	fmt.Println(g)

	var(
		age1 = 100
		age2 = 200
	)                 // 如果是多种类型 也可以使用集合
	age1, age2 = age2, age1  //可以轻松实现交换变量的值
	fmt.Println(age1, age2)
	fmt.Println("---------------------------------------------------")
	/**
	常量
	*/

}
