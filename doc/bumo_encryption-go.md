# BUMO GO ENCRYPTION使用文档

## 概述
本文档简要概述Bumo Go 部分常用接口文档



- [包引用](#包引用)
- [生成地址](#生成地址)
	- [获取编码后私钥，公钥，地址](#获取编码后私钥，公钥，地址)
    - [私钥生成公钥](#私钥生成公钥)
    - [公钥生成地址](#公钥生成地址)
    - [私钥是否有效](#私钥是否有效)
	- [公钥是否有效](#公钥是否有效)
	- [地址是否有效](#地址是否有效)

- [签名](#签名)
	- [签名](#签名)
	- [验签](#验签)


用于生成公私钥和地址，以及签名，和验签，只支持ED25519。

## 包引用
所依赖的golang包在src文件夹中寻找，依赖的golang包如下：

    go get github.com/bumoproject/bumo-encryption-go//获取包
```
//底层依赖
1. "github.com/bumoproject/bumo-encryption-go/src/bumo/base58"
2. "github.com/bumoproject/bumo-encryption-go/src/3rd/ed25519/edwards25519"
//示例依赖
3. "github.com/bumoproject/bumo-encryption-go/src/keypair"
4. "github.com/bumoproject/bumo-encryption-go/src/signature"
```

## 生成地址
### 获取编码后私钥，公钥，地址
调用keypair包里的Create方法

示例如下：
```go
 public, private, address, err := keypair.Create()
	if err != nil {
		return
	}
	fmt.Println("public key:", public)
	fmt.Println("private key:", private)
	fmt.Println("address:", address)
	
	
// 	public key: b001c0f7b59e209bedf66dc2966d70d5a98ebb97ccb7351048363f75542e90ae9722198f0189
//    private key: privbvjguD3bQvWhvcH472JPb6U1aSwELbHrYiGA6JnMiBKdd76NqjPP
//    address: buQhen4Kme3wRweCpCpPsZSe4dALVEtVwv2Q
```


###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
public     |  String        |  账户区块链公钥     |
private     |  String        |  账户区块链私钥     |
address     |  String        |  账户区块链地址     |
err     |  error        |  错误    |


### 私钥生成公钥


示例如下：
```go
    private := "privbvjguD3bQvWhvcH472JPb6U1aSwELbHrYiGA6JnMiBKdd76NqjPP"
	pubFromPriv,err := keypair.GetEncPublicKey(private)
	if err != nil {
		return
	}
	fmt.Println("public key from private key: ", pubFromPriv)
	
//	public key from private key:  b001c0f7b59e209bedf66dc2966d70d5a98ebb97ccb7351048363f75542e90ae9722198f0189
	
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
private |  String   | 账户区块链私钥 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
pubFromPriv     |  String        |  账户区块链公钥     |
err     |  error        |  错误    |






### 公钥生成地址


示例如下：
```go
    public := "b001c0f7b59e209bedf66dc2966d70d5a98ebb97ccb7351048363f75542e90ae9722198f0189"
	addrFromPub,err := keypair.GetEncAddress(public)
	if err != nil {
		return
	}
	fmt.Println("address from public key: ", addrFromPub)
	
	//address from public key:  buQhen4Kme3wRweCpCpPsZSe4dALVEtVwv2Q
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
public |  String   | 账户区块链公钥 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
addrFromPub     |  String        |  账户区块链地址    |
err     |  error        |  错误    |



### 私钥是否有效


示例如下：
```go
    private := "privbvjguD3bQvWhvcH472JPb6U1aSwELbHrYiGA6JnMiBKdd76NqjPP"
	privatestatus := keypair.CheckPrivateKey(private)
	fmt.Println("private status: ", privatestatus)
	
	//private status:  true
	
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
private |  String   | 账户区块链私钥 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
privatestatus     |  bool        |  私钥是否有效    |



### 公钥是否有效


示例如下：
```go
    public := "b001c0f7b59e209bedf66dc2966d70d5a98ebb97ccb7351048363f75542e90ae9722198f0189"
    publicstatus := keypair.CheckPublicKey(public)
	fmt.Println("address status: ", publicstatus)
	
	//public status:  true
	
	
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
public |  String   | 账户区块链公钥 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
publicstatus     |  bool        |  公钥是否有效    |




### 地址是否有效


示例如下：
```go

    address := "buQhen4Kme3wRweCpCpPsZSe4dALVEtVwv2Q"
	addressstatus := keypair.CheckAddress(address)
	fmt.Println("address status: ", addressstatus)
	
	
	//address status:  true
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
address |  String   | 账户区块链地址 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
addressstatus     |  bool        |  地址是否有效    |



## 签名
### 签名

示例如下：
```go
	var message []byte
	private := "privbvjguD3bQvWhvcH472JPb6U1aSwELbHrYiGA6JnMiBKdd76NqjPP"
	sign, err := signature.Sign(private, message)
	if err != nil {
		return
	}
	fmt.Println("sign: ", sign)
	
	//sign:  d66c8f6f008e07e8b825ef7c98092bb5082e03f7b910a7c103dff7ac6e482a10baf2782efe4010f1ebf3a56d06d0542abd267645604aa31ed926b9ad3ba8140a

```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
private |  String   | 账户区块链私钥 |
message |  String   | 签名 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
sign     |  string       |  签名后数据    |
err     |  error       |  错误    |



### 验签

示例如下：
```go

    sign := "d66c8f6f008e07e8b825ef7c98092bb5082e03f7b910a7c103dff7ac6e482a10baf2782efe4010f1ebf3a56d06d0542abd267645604aa31ed926b9ad3ba8140a"
    public := "b001c0f7b59e209bedf66dc2966d70d5a98ebb97ccb7351048363f75542e90ae9722198f0189"
	verify := signature.Verify(public, message, sign)
	fmt.Println("verify: ", verify)
	
	
	//verify:  true
```

###### 入参
参数    |   类型    |      描述      |
------- | --------- | -------------- |
public |  String   | 账户区块链公钥 |
message |  String   | 需要签名信息 |
message |  String   | 签名 |
###### 返回参数
参数        |   类型         |       描述          |       
----------- | -------------- | ------------------- |
verify     |  bool       |  是否验证成功    |

