# httproxy
A simple HTTP/HTTPS/SOCKS5 proxy.

# Installation

`go build`

# Usage

* Install httproxy on the proxy machine and start it with `httproxy -addr=:8080`
* Can work with [KCPTUN][2] to build a tunnel http/socks5 proxy...
* Test with cURL
  ```sh
  # https proxy
  curl -x 127.0.0.1:8080 https://github.com

  # socks5 proxy
  curl --socks5-hostname 127.0.0.1:8080 http://github.com
  ```

# More
