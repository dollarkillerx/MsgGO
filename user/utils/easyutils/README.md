# easyutils
easyutils Golang 常用工具库
``` 
.
├── crypto.go  加密解密相关
├── go.mod
├── LICENSE
├── README.md
├── simpleTime.go  时间相关
├── tootl_test.go 测试 
└── uuid.go  uuid
```

把生活和工作当中常用的东西抽离出来
严格单元测试 保证代码质量
献丑了

### 简答用法
- 把str转为md5
    ``` 
    md5str := Md5Encode(string)
    ```
- 获取当前时间戳,时区默认亚洲上海
    ``` 
    timeString := GetCurrentTime()
    ```
- 时间戳转换为日期
    ``` 
    日期,err := GetTimeToString(时间戳string)
    ```
- 日期转换为时间戳
    ``` 
    时间戳str,err := GetTimeStringToTime(日期)
    ```
- 获取uuid
    ``` 
    uuidstr,err := NewUUID()
    ```
- 获取当前uuid不带-
    ``` 
    uuidstr,err := NewUUIDSimplicity()
    ```