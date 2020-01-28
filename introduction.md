# 写在前面

### 没有代码实现的数据结构和算法都是耍流氓

对于码农来说，只是理解数据结构与算法的思想，或者只写些伪代码，而不用代码去具体实现一遍，那与纸上谈兵无异。

### 工欲善其事，必先利其器

既然决定要撸数据结构与算法的代码，编程语言的选择至关重要，就如刽子手的刀，一个优秀的刽子手，只有配上一把趁手的鬼头刀，那样才会砍得干净利落。

诚然，C语言是撸数据结构与算法最好的编程语言，它能让编写者最大程度地理解数据结构与算法的核心，可是C语言中严格的先声明后调用(特指函数)以及内存的动态分配和回收，给普通码农增加了不少负担。

那么有没有更好的选择呢？

Golang可能就是那把万能的鬼头刀，可长可短，还带自动清洗功能，使每个刽子手都用得趁手。

首先，Golang中有指针，虽然指针不是万能的，但对于数据结构来说，没有指针是万万不能的。在实现数据结构的代码中，不使用指针就与其思想相去甚远，这样的代码是没有灵魂的。

其次，Golang借鉴了许多语言的特性，对拥有其他编程语言基础的码农十分友好。
1. 垃圾自动回收机制 大概是借鉴java
2. var声明变量以及多返回值 大概是借鉴js
3. 切片 大概是借鉴python
4. switch支持字符串 大概是借鉴php
5. 结构体、类型定义和指针 大概是借鉴c
……

最后，如果看了首先和其次，还不选择Golang实现，那还有天理吗，还有法律吗？