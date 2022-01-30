# chapter2-异常处理
## 作业内容
1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 思考分析

### 建库建表
本示例依赖的库表结构
```mysql
CREATE SCHEMA go_advanced_tc;

CREATE TABLE `student` (
`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
`name` varchar(32) NOT NULL DEFAULT '' COMMENT '姓名',
`age` int(11) NOT NULL COMMENT '年龄',
PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='学生表';
```
### 代码结构
```text
.
├── README.md
├── go.mod
├── go.sum
├── main.go
├── src
│   ├── mysql
│   │   └── mysql.go
│   └── student
│       └── student.go
└── vendor

```
> 数据库的连接方式、库名，都硬编码在main.go中，如有不同则手动更改
> 
> dsn := "root:Root1234@tcp(10.69.77.221:8306)/go_advanced_tc?charset=utf8mb4"


- main.go 是程序主入口
- src 程序核心代码
  - mysql
    - mysql.go 用于创建和关闭数据库连接
  - student
    - student.go 基于student表实现的model和dao
- vendor 不包含在代码库中

### 运行方式

```shell
# 下载依赖包并集成到工程内
go mod tidy
go mod vendor

# 直接运行
go run main.go
```

### 编码思路

#### mysql连接与关闭

在mysql/mysql.go中，创建dbRepo结构体和DbRepo接口，并通过New函数返回DbRepo对象。

在New函数中通过`sql.Open("mysql", dsn)`对数据库进行建链，并通过`db.Ping()`对数据库进行探活。

当发生err时，直接`panic(err)`使进程终止。

> 在这里考虑到程序强依赖mysql，如果无法连接mysql，则直接panic，不需要再特殊处理或抛给上层。

#### student

在student/student.go中，创建model结构体表示student实体，dao结构体和Dao接口用来实现持久化操作。

Dao接口中实现`QueryAll()`和`QueryRow()`，分别调用`sql.DB`的`Query()`和`QueryRow()`

**以`QueryAll()`为例说明errors.wrap()的必要性**

在最初的编码中没有使用wrap，如下
```go
func (d dao) QueryAll() ([]*model, error) {
  sql := "select * from student"
  rows, err := d.dbRepo.GetDb().Query(sql, 1)
  if err != nil {
    return nil, err
  }
  defer rows.Close()
  
  // other codes
}
```

上述代码存在明显的问题，在第3行调用Query时多传了一个参数，在运行main.go时抛错
```text
new db repo success
sql: expected 0 arguments, got 1
```

尴尬的要死，并不知道是哪里报的错，只能一层层查调用

那么当加入wrap后，代码和报错如下
```go
func (d dao) QueryAll() ([]*model, error) {
  sql := "select * from student"
  rows, err := d.dbRepo.GetDb().Query(sql)
  if err != nil {
    return nil, errors.Wrap(err, "query err")
  }
  defer rows.Close()
}
```
```text
new db repo success
sql: expected 0 arguments, got 1
query err
chapter02error/src/student.dao.QueryAll
        /Users/chkinglee/coding/github/chkinglee/geektime/go_advanced_tc/chapter02error/src/student/student.go:40
main.main
        /Users/chkinglee/coding/github/chkinglee/geektime/go_advanced_tc/chapter02error/main.go:27
runtime.main
        /Users/chkinglee/.g/go/src/runtime/proc.go:255
runtime.goexit
        /Users/chkinglee/.g/go/src/runtime/asm_arm64.s:1133
```

wrap包装了错误，携带了err的堆栈信息，定位问题简单清晰

**以`QueryRow()`为例回答作业问题**

`sql.ErrNoRows`在调用`sql.DB`的`QueryRow()`时可能会发生，但这种情况对于业务来说是正常的，就是查无此人嘛。

那么在持久层，查询数据时遇到了`sql.ErrNoRows`，可以通过判断err的值（这里使用了`errors.Is()`）来处理

> 在这里考虑到如果持久层发生`sql.ErrNoRows`，则直接`return nil, nil`即可
> 
> 因为在业务层要对返回的student做处理

```go
	stuOne, err := studentDao.QueryRow()
	if err != nil {
		fmt.Printf("err: %+v",err)
		return
	}
	if stuOne == nil {
		fmt.Println("stu not found")
		return
	}
	stuOne.ToString()
```

以上代码相当于业务层逻辑，首先要处理err，然后是对stuOne的业务处理，当然，也可以直接将`if stuOne == nil`的代码块放到`ToString()`的实现中，使业务层代码更简洁

最终代码

```go
func (m *model) ToString() {
	if m == nil {
		fmt.Println("stu not found")
		return
	}
	fmt.Println(fmt.Sprintf("Id: %d, Name: %s, Age: %d", m.Id, m.Name, m.Age))
}
```
```go
	stuOne, err := studentDao.QueryRow()
	if err != nil {
		fmt.Printf("err: %+v",err)
		return
	}
	stuOne.ToString()
```