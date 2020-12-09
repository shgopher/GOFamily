# 很早之前写的summary，不扔了吧 留个纪念。

## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)


    > # linux 的一些技术总结（详细的不说了，你们可以自己去看书，我只写大概的东西）

    #### 1输入
    使用echo，不使用‘’或者“”，只要记得如果使用特殊符号的时候使用\即可，
    >如果想了解详细的例如使用‘’，和使用“”，请自行查阅资料

    ```
    echo hello world\!
    // hello world !
    ```
    ### 2查看环境变量
    使用 cat /proc/$pid/environ
    >这个$pid可以使用 pgrep +进程名字获取

    ```
    pgrep setting
    // 10248
    cat /proc/10248/environ
    //环境变量，有一大堆！

    ```
    >返回值是以name=value这种形式来显示并且分隔符是null,那么如果想更换就可以使用tr命令， 例如

    ```
    cat /proc/10248/environ |tr '\0' '\n'
    ```
    结果就是以换行进行输出
    ![this is a picture](./picture/picture1.png)

## [目录](.https://github.com/shgopher/GOFamily/tree/master/%E5%85%A5%E9%97%A8%E7%AF%87/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F/shell)


