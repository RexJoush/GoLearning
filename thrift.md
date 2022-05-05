## IDL 详解

* Thrift 是一个典型的 CS 结构，客户端和服务端可以使用不同的语言开发，既然客户端和服务端能使用不同的一眼开发，那么就需要一种中间语言来关联客户端语言和服务端语言，这种语言就是 IDL（Interface Description Language）

* Thrit 采用 IDL 来定义通用的服务接口，然后通过 Thrift 提供的编译器，可以将服务接口编译成不同语言编写的代码，通过这个方式来实现跨语言的功能

## IDL 语法

#### 基本类型（Base Types）

* 基本类型：不管那一种语言，都支持的数据形式，Apache Thrift 支持以下几种

  | Type   | Desc              | Java             | Go      |
  | ------ | ----------------- | ---------------- | ------- |
  | i8     | 有符号8位整数     | byte             | int8    |
  | i16    | 有符号16位整数    | short            | int16   |
  | i32    | 有符号32位整数    | int              | int32   |
  | i64    | 有符号64位整数    | long             | int64   |
  | double | 64位浮点数        | double           | float64 |
  | bool   | 布尔值            | boolean          | bool    |
  | string | 字符串，utf-8编码 | java.lang.String | string  |

#### 特殊类型（Special Types）

* binary，为编码的字节序列，是 string 的一种特殊形式，这种类型主要是方便某些场景下 java 调用，对应的是 `java.nio.ByteBuffer`类型，Go 中是 `[]byte`

#### 集合容器（containers）

* 使用容器类型必须指定泛型，否则无法编译 idl 文件，其次，泛型中的基本类型，Java 语言中会被替换为对应的包装类型。集合中的勇士可以是除了 service 之外的任何类型，包括 execption

    | Type      | Desc                           | Java           | Go                                 |
    | --------- | ------------------------------ | -------------- | ---------------------------------- |
    | list<T>   | 元素有序，允许重复             | java.util.List | []T                                |
    | set<T>    | 元素无序，不允许重复           | java.util.Set  | []T,Go中没有  set 集合，以数组代替 |
    | map<K, V> | key-value数据结构，key不许重复 | java.util.Map  | map[K] V                           |
    
    ```thrift
    struct Test {
    	1: map<string, User> usermap
    	2: set<i32> intset,
    	3: list<double> doublelist
    }
    ```

#### 常量及类型别名（Const && Typeof）

```thrift
// 常量定义
const i32 MALE_INT = 1
const map<i32, string> GENDER_MAP = {1: "male", 2: "female"}
// 某些数据类型比较长可以用别名简化
typeof map<i32, string> gmp
```

#### struct 类型

* 在面向对象语言中，表现为类定义，在弱类型语言、动态语言中，表现为结构/结构体。定义格式如下

  ```thrift
  struct <结构体名称> {
  	<序号>:[字段性质] <字段类型> <字段名称> [= <默认值>] [;|,]
  }
  ```

* 例如

  ```thrift
  struct User {
  	1: required string name, // 该字段必须填写
  	2: optional i32 age = 0; // 默认值，且不必须传
  	3: bool gender // 默认字段类型为 optional
  }
  
  struct bean {
  	1: i32 number = 0,
  	2: i64 bigNumber,
  	3: double decimals,
  	4 string name = "thrift"
  }
  ```

* struct 有以下一些约束
  * struct 不能继承，但可以嵌套，不能嵌套自己
  * 其成员必须都有明确的类型
  * 成员是被正整数编号过的，其中的编号不能重复，为了在传输过程编码使用
  * 成员分隔符可以是（ , ）也可以是（ ; ），而且可以混用
  * 字段会有 optional 和 required 之分，和 protobuf 一样，如果不指定则为无类型，可以不填充该值，但在序列化的传输过程中，也会序列化。opional 是不填充则不序列化，required 是必须填充也必须序列化
  * 每个字段可以设置默认值
  * 同一文件可以定义多个 struct，也可以定义在不同文件，进行 include 引入

#### 枚举（enum）

* Thrift 不支持枚举类，枚举常量必须是32为的正整数

  ```thrift
  enum HttpStatus {
  	OK = 200,
  	NOTFOUND = 404
  }
  ```

#### 异常（Exceptions）

* 异常在语法和功能上类似结构体，差别是异常使用关键字 exception，而异常是继承每种语言的基础异常类

  ```thrift
  exception MyException {
  	1: i32 errorCode,
  	2: string message
  }
  
  // 异常的使用方法
  service ExampleService {
  	string GetName() throws (1: myException e1, 2: myException e2)
  }
  ```

#### service（服务定义类型）

* 服务的定义方法在语义上等同于面向对象语言中的接口

  ```thrift
  service HelloService {
  	i32 satIny(1: i32 param)
  	string sayString(1: string param)
  	bool seyBoolean(1: bool param)
  	void sayVoid()
  }
  ```

* 上面编译后的 Java 代码

  ```java
  public class HelloService {
      public interface Idace {
          public int sayInt(int param) throws org.apache.thrift.TException
          public java.lang.String sayString(java.lang.String param) throws org.apache.thrift.TException
          public boolean sayBoolean(boolean param) throws org.apache.thrift.TException
          public void sayVoid() throws org.apache.thrift.TException
      }
      // ...
  }
  ```

#### Namespace（命名空间）

* Thrift 中的命名空间类似于 C++ 中的 namespace 和 java 中的 package，它们提供了一种组织（隔离）代码的渐变方式，命名空间也可以用于解决类型定义中的名字冲突。

* 由于每种语言均有自己的命名空间定义方式，thrift 允许开发者针对特定语言定义 namespace

  ```thrift
  namespace java com.example.test
  // 转换成
  package com.example.test
  ```

* 例如：

  ```thrift
  namespace java com.joush // 命名空间定义，规范 namespace + 语言 + 包路径
  service Hello { // 接口定义
  	string getWoid(), // 定义方法
  	void writeWorld(1: string words) // 参数类型指定
  }
  ```

#### Comment（注释）

* Thrift 支持 C 多行风格和 Java/C++ 单行风格的注释

#### Include

* 便于管理、重用和提高模块性/组织性，常常分割 thrift 定义在不同的文件中，包含文件搜索方式与 C++ 一样，thrift 允许文件包含其他 thrift 文件，用户需要使用 thrift 文件名作为前缀访问被包含的对象

* thrift 文件名需要使用双引号包含，且末尾没有逗号或分号

  ```thrift
  include "test.thrift"
  // ...
  struct StSearchResult {
  	1: i32 uid
  	// ...
  }
  ```

  
