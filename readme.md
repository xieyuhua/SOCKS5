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

```
// 'function FindProxyForURL(url, host) {return "HTTPS 223.xyz.xyz;HTTPS beitai.520spciex.org;";}',   同时支持http和https
// 'function FindProxyForURL(url, host) {if ( shExpMatch(host, "*github*") ){	return "HTTPS 223.xyz.xyz;";}else{return "DIRECT";}}',
// "SOCKS5 127.0.0.1:1080;PROXY 127.0.0.1:8888;HTTPS 223.xyz.xyz;DIRECT"
cc = 'function FindProxyForURL(url, host) {return "PROXY 116.205.229.85:80;PROXY 124.261.0.127:8080;HTTPS 223.xyz.xyz;";}',
startProxy = function() {
	chrome.proxy.settings.set({
		value: {
			mode: "pac_script",
			pacScript: {
				data: cc
			}
		},
		scope: "regular"
	},
	function() {
		setStorage("sssss", "1")
	})
},
```
