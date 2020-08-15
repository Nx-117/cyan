# Cyan

####Cyan是目前自己常用的一些方法的集合。

####我是想把Cyan做为一个长期维护的项目，随着时间的推移会更新越来越多的常用代码。

# 文档
#### 以下仅列出本项目所有的函数名字及作用，具体使用方法可以参考代码注释
+ captcha
    + GenerateCaptcha   (生成验证码图片Base64)
+ config
    + LoadConfig    (读取配置文件)
+ date
    * Now (获取当前时间)
    * NowByFormat (根据指定格式获取当前时间)
    * NowByTime (根据传入的time获取日期)
    * GetMillisecond (获取当前毫秒值)
    * NowByTimeAndFormat (根据传入的time和指定格式获取日期)
    * TimestampFormatToString (根据时间戳和指定时间格式转换字符串)
    * TimestampDefaultFormatToString (根据时间戳转换字符串)
    * Week (获取当前是第几周)
    * Weekday (获取当天日期是周几)
    * DateYMD (获取年、月、日)
    * Tomorrow (获取明天日期)
    * Yesterday (获取昨天日期)
    * GetSpecifiedDateByYMD (根据指定天数获取日期)
    * GetSpecifiedDateByYMDAndFormat (根据指定天数和指定格式,获取日期)
    * StringConverToTime (根据时间字符串转换成time类)
+ encryption
    * Md5String (MD5加密)
    * Hash (md5、sha1、sha256、sha512加密算法)
    * AESEncrypter (aes加密算法)
    * AESDecrypter (aes解密算法)
    * GenerateKey (生成公钥、私钥)
    * RSAEncrypter (RSA加密算法)
    * RSADecrypter (RSA解密算法)
    * GetKeyStringByFile (密钥证书转换成字符串)
+ errors
    * New (生成自定义error信息)
+ file
    * Read (读取文件内容)
    * ReadToString (读取文件到字符串)
    * CheckFileIsExist (判断文件是否存在)
+ file
    * Get (get请求方式)
    * GetByParams (get请求并携带参数)
    * GetByParamsAndHeads (get请求并携带参数和请求头)
    * PostForm (post请求form表单方式)
    * PostFormHeads (post请求form方式并设置head)
    * PostJson (post请求json方式)
    * PostJsonHead (post请求json方式并设置head)
    * PostSendFileAndHead (post请求发送文件并携带head)
+ jwt
    * GenerateToken (创建token)
    * ParseToken (解析token)
+ log
    * IntiLog (初始化日志文件,搭配gin框架使用)
+ randMath
    * RandMath4 (生成4位随机数) 
    * RandMath6 (生成6位随机数) 
    * RandMath (生成 0至age 指定范围的随机数) 
+ task
    * StartJob (定时任务)
    * Timing (定时任务)
