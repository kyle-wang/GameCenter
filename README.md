# GameCenter Login Verify

## 说明

> Apple GameCenter的登录服务器验证golang版本.
>
> 根据Apple的官方文档实现，文档地址：https://developer.apple.com/documentation/gamekit/gklocalplayer/1515407-generateidentityverificationsign.
>
> 服务器接受的参数为：
>
>> puk : public key 用于下载public key的地址，会进行缓存
>
>> data : 客户端根据文档第6，7步生成的数据，通过base64加密后传递给服务器
>
>> sig : 客户端传递给服务器的Apple返回的签名数据，也是经过base64加密后的数据
>
>服务器返回参数：
>
>> error : 如果返回参数为nil，则验证通过，本次登录合法，否则为本次登录非法

## 使用示例

> 		func verifyGameCenter(username, data, sig, puk string) (error, string) {
>			if err := GameCenter.Verify(puk, data, sig); err == nil {
>				return nil, username
>			} else {
>				return err, ""
>			}
>		}
>

