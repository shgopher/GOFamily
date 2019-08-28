
```go

public static int rand7() {
    int v = rand5() * 5 + rand5();
    if (v < 21) {
        return v % 7;
    }
    return rand7();
}

```
