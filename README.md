# colume-search
一个纯go编写的mini型倒排索引工具。项目名暂时这样叫，后面心血来潮了会改。
##基本思想
本项目单纯是为了解决检索某个数据而提供的一种解决办法，自身并不承担类似mysql或者其他DB一样存储关系型数据的功能，仅仅是提供了一种能够高效检索数据的办法。本项目认为，任何具备检索逻辑的数据在结构上都具备以下的特性：
	
	{
	    "key":1,  //某种关系型数据的数据主键
	    "data":{
	        "attr1":"value1",  //该条记录所有需要检索的index和value
	        "attr2":"value2",
	        "attr3":"value3",
	        "attr4":"value4",
	        "attr5":"value5",
	    }
	}
	
检索到数据之后，自行再结合其他DB查询出对应的数据(后续有精力了再考虑搞这个)，项目提供了web服务，如果不想使用web服务您也可以自行修改代码，作为内嵌的检索引擎使用。


# 编译

    git clone https://github.com/hide-in-code/colume-search.git

    cd colume-search

    go build -o colume-search main.go

# 使用

    ./colume-search
执行之后程序会启动一个web服务，服务监听本地`5200`端口，可以通过web协议进行数据的添加，查询和修改，接口规范符合`restful`风格API


### 新增数据
	POST 127.0.0.1:5200/{{table}}  //暂时叫table，没有想到好的名字
	
	inputdata
	{
	    "key":2,   //查询结果关联的主键key
	    "data":{
	        "attr1":"value11",  //每一个记录的index和value，倒排之后可以根据的index和value 查询出 对应记录的 key
	        "attr2":"value22",
	        "attr3":"value33",
	        "attr4":"value44",
	        "attr5":"value55",
	    }
	}

### 根据index和value检索某一条数据(key)
	GET 127.0.0.1:5200/{{table}}?index=index1&value=value1

### 检索所有唯一index
	GET 127.0.0.1:5200/{{table}}/allindex
	
### 检索所有index值对应的数量
	GET 127.0.0.1:5200/{{table}}/allindexcount

### 项目依赖
- [gin框架](https://github.com/gin-gonic/gin)(也可以不使用web服务，单纯作为内嵌的检索库)

### TODO
- 数据修改

# 协议
MIT License