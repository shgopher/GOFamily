## 模版
模版模式的核心是，公有方法 和子方法的解耦。

也就是说，一个接口中的众多方法，分为公有方法和子方法，公有方法是事先定义好的，强制子方法去实现各自的方法。

```go
func main(){
	a1 := new(A1)	
	Do(a1)
	// 开始第二个部分的执行。
	a2 := new(A2)	
	Do(a2)
}

type Cook interface {
	first()
	second()
	last()
}

type A struct {
}

func (A) first()  {
	fmt.Println("first")
}

type A1 struct {
	A
}

func (a *A1) second()  {
	fmt.Println("second1")
}
func (a *A1) last()  {
	fmt.Println("last1")
}
type A2 struct {
	A
}
func (a *A2) second()  {
	fmt.Println("second2")
}
func (a *A2) last()  {
	fmt.Println("last2")
}

func Do(cook Cook){
	cook.first()
	cook.second()
	cook.last()
}
```