一个完整的可用的beego架构，除了tests
```
User registration:
request:
   POST  /v1/user
    {
        name: wujq,
        passwd: 123456
        
    }
response:
    {
        status: "success"  #  or failed
        
    }
```
``` 
User login:
   request:
        POST /v1/login
        {
            name: wujq,
            passwd: 123456  
        }
    response:
        {
            status: failed,
            "ermsg": username/password is wrong
        } 
   
        or 
        {
            status: success,
            token: ...
        }

```
```
User reflashToken:
    request:
        Post /v1/reflash
    response:
        {
            status: failed
        } 
   
        or 
        {
            status: success,
            token: ...
        }
        
``` 
