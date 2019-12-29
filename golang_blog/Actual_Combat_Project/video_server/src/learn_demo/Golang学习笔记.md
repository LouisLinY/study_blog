# Golang学习笔记

## Go语言之Json与MD5

### json简介

#### JSON   GO语言内置的encoding/json标准库（性能不是很好）

####              github.com/pquerna/ffjson/ffjson



#### json编码：

```go
func Marshal(v interface{}) ([]byte, error)
```

#### json解码：

```go
func Unmarshal(data []byte, v interface{}) error
```

​              

### MD5   Go语言内置的crypto/md5标准库

```go
MD5Inst := md5.New()
MD5Inst.Write([]byte("test md5"))
Result := MD5Inst.Sum([]byte[""])
fmt.Println("%x\n\n", Result)
```





sql.Open()函数用来打开一个注册过的数据库驱动，Go-MySQL-Driver中注册了mysql这个数据库驱动，第二个参数是DNS(Data source name),它是Go-MySQL-Driver定义的一些数据库链接和配置信息。格式如下：

```go
user@unix/path/to/socket/dbname?charset=utf8
user:password@tcp(localhost:5555)/dbname?charset=utf8

```

