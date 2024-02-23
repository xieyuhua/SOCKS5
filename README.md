# socks5
socks5 proxy server builds with go/golang

### Features
* Both TCP and UDP supported
* "No Auth" authentication supported
* User/Password authentication supported

## Install
```
$ go mod tidy
$ go build
```

## Usage
* -p
  * port on listen (default 1080)
* -user (option)
  * username
* -pwd (option)
  * password
        

```
$ curl -x socks5h://192.168.5.254:1080  www.baidu.com
$ curl -x socks5h://username:password@192.168.5.254:1080  www.baidu.com
```
