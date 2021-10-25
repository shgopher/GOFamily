# encoding
encoding定义了从字节到文本的转译。

此包拥有的众多子包，比如常见的base64， json 等都属于转码操作

其中包括了

- ascii85
- asn1
- base32
- base64
- binary
- csv
- gob
- hex
- json
- pem
- xml

encoding包定义了四个接口类型，分别是 `type BinaryMarshaler` , `type BinaryUnmarshaler`, `type TextMarshaler`, `type TextUnmarshaler`

根据go的面向接口编程的理念，encoding定义的抽象接口，在各个子包中，都会调用，比如说gob，json都会检查这些接口的实现，意思就是说，你如果自己实现了这些接口，那么子包就会调用这些接口。

面向接口编程的例子，例如json包：
```go
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
```

`json.Marshal`函数会递归的处理值。如果一个值实现了`Marshaler`接口切非nil指针，会调用其`MarshalJSON`方法来生成json编码。nil指针异常并不是严格必需的，但会模拟与`UnmarshalJSON`的行为类似的必需的异常。

但是如果你没有实现这个接口，那么go的标准库就会自己定义一套东西来去帮你做这套东西，所以go的原则就是“给你选择权，但是你不做的话，我帮你做”。