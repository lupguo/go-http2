## Test Go HTTP/2 Server Push

### Cert Generate
```
openssl req -nodes -x509 -newkey rsa:1024 -keyout gohttp2.key -out gohttp2.cert -days 3650 \ 
    -subj "/C=CN/ST=GD/L=SZ/O=HQYG, Inc./OU=Tech./CN=TK.CA/emailAddress=tkstorm1988@gmail.com"
```

### Running
```
// clone
git clone https://github.com/tkstorm/go-http2
// generate cert
openssl req -nodes -x509 -newkey rsa:1024 -keyout gohttp2.key -out gohttp2.cert -days 3650 -subj "/C=CN"
// run web server
go run *.go
```

### Set push content

```
// set http2 push
func h2push(w http.ResponseWriter, s ...string) {
	if p, ok := w.(http.Pusher); ok {
		for _, t := range s {
			p.Push(t, nil)
		}
	}
}
```

### Result
<img src="/img/h2push.jpg" alt="h2push" width="500">

### Nghttp debug
```
$ nghttp -unv https://localhost:2345
[  0.017] Connected
[WARNING] Certificate verification failed: Hostname mismatch
The negotiated protocol: h2
[  0.023] recv SETTINGS frame <length=24, flags=0x00, stream_id=0>
          (niv=4)
          [SETTINGS_MAX_FRAME_SIZE(0x05):1048576]
          [SETTINGS_MAX_CONCURRENT_STREAMS(0x03):250]
          [SETTINGS_MAX_HEADER_LIST_SIZE(0x06):1048896]
          [SETTINGS_INITIAL_WINDOW_SIZE(0x04):1048576]
[  0.023] send SETTINGS frame <length=12, flags=0x00, stream_id=0>
          (niv=2)
          [SETTINGS_MAX_CONCURRENT_STREAMS(0x03):100]
          [SETTINGS_INITIAL_WINDOW_SIZE(0x04):65535]
[  0.023] send SETTINGS frame <length=0, flags=0x01, stream_id=0>
          ; ACK
          (niv=0)
[  0.023] send PRIORITY frame <length=5, flags=0x00, stream_id=3>
          (dep_stream_id=0, weight=201, exclusive=0)
[  0.023] send PRIORITY frame <length=5, flags=0x00, stream_id=5>
          (dep_stream_id=0, weight=101, exclusive=0)
[  0.023] send PRIORITY frame <length=5, flags=0x00, stream_id=7>
          (dep_stream_id=0, weight=1, exclusive=0)
[  0.023] send PRIORITY frame <length=5, flags=0x00, stream_id=9>
          (dep_stream_id=7, weight=1, exclusive=0)
[  0.023] send PRIORITY frame <length=5, flags=0x00, stream_id=11>
          (dep_stream_id=3, weight=1, exclusive=0)
[  0.023] send HEADERS frame <length=38, flags=0x25, stream_id=13>
          ; END_STREAM | END_HEADERS | PRIORITY
          (padlen=0, dep_stream_id=11, weight=16, exclusive=0)
          ; Open new stream
          :method: GET
          :path: /
          :scheme: https
          :authority: localhost:2345
          accept: */*
          accept-encoding: gzip, deflate
          user-agent: nghttp2/1.39.1
[  0.026] recv WINDOW_UPDATE frame <length=4, flags=0x00, stream_id=0>
          (window_size_increment=983041)
[  0.026] recv SETTINGS frame <length=0, flags=0x01, stream_id=0>
          ; ACK
          (niv=0)
[  0.026] recv (stream_id=13) :method: GET
[  0.026] recv (stream_id=13) :scheme: https
[  0.027] recv (stream_id=13) :authority: localhost:2345
[  0.027] recv (stream_id=13) :path: /css/style.css
[  0.027] recv PUSH_PROMISE frame <length=30, flags=0x04, stream_id=13>
          ; END_HEADERS
          (padlen=0, promised_stream_id=2)
[  0.028] recv (stream_id=13) :method: GET
[  0.028] recv (stream_id=13) :scheme: https
[  0.028] recv (stream_id=13) :authority: localhost:2345
[  0.028] recv (stream_id=13) :path: /img/blog-red-logo.jpg
[  0.028] recv PUSH_PROMISE frame <length=25, flags=0x04, stream_id=13>
          ; END_HEADERS
          (padlen=0, promised_stream_id=4)
[  0.028] recv (stream_id=4) :status: 200
[  0.028] recv (stream_id=4) content-type: image/jpeg
[  0.028] recv (stream_id=4) date: Thu, 19 Sep 2019 08:14:02 GMT
[  0.028] recv HEADERS frame <length=35, flags=0x04, stream_id=4>
          ; END_HEADERS
          (padlen=0)
          ; First push response header
[  0.028] recv (stream_id=2) :status: 200
[  0.028] recv (stream_id=2) content-type: text/css
[  0.029] recv (stream_id=2) content-length: 56
[  0.030] recv (stream_id=2) date: Thu, 19 Sep 2019 08:14:02 GMT
[  0.030] recv HEADERS frame <length=14, flags=0x04, stream_id=2>
          ; END_HEADERS
          (padlen=0)
          ; First push response header
[  0.030] recv DATA frame <length=16384, flags=0x00, stream_id=4>
[  0.030] recv (stream_id=13) :status: 200
[  0.030] recv (stream_id=13) content-type: text/html
[  0.030] recv (stream_id=13) content-length: 258
[  0.030] recv (stream_id=13) date: Thu, 19 Sep 2019 08:14:02 GMT
[  0.030] recv HEADERS frame <length=16, flags=0x04, stream_id=13>
          ; END_HEADERS
          (padlen=0)
          ; First response header
[  0.030] recv DATA frame <length=16384, flags=0x00, stream_id=4>
[  0.030] recv DATA frame <length=16384, flags=0x00, stream_id=4>
[  0.030] recv DATA frame <length=16383, flags=0x00, stream_id=4>
[  0.030] send WINDOW_UPDATE frame <length=4, flags=0x00, stream_id=0>
          (window_size_increment=32768)
[  0.030] send WINDOW_UPDATE frame <length=4, flags=0x00, stream_id=4>
          (window_size_increment=32768)
[  0.030] send WINDOW_UPDATE frame <length=4, flags=0x00, stream_id=0>
          (window_size_increment=32767)
[  0.030] send WINDOW_UPDATE frame <length=4, flags=0x00, stream_id=4>
          (window_size_increment=32767)
[  0.031] recv DATA frame <length=258, flags=0x01, stream_id=13>
          ; END_STREAM
[  0.031] recv DATA frame <length=56, flags=0x01, stream_id=2>
          ; END_STREAM
[  0.031] recv DATA frame <length=16384, flags=0x00, stream_id=4>
[  0.031] recv DATA frame <length=16384, flags=0x00, stream_id=4>
[  0.031] recv DATA frame <length=3152, flags=0x00, stream_id=4>
[  0.031] recv DATA frame <length=0, flags=0x01, stream_id=4>
          ; END_STREAM
[  0.031] send GOAWAY frame <length=8, flags=0x00, stream_id=0>
          (last_stream_id=4, error_code=NO_ERROR(0x00), opaque_data(0)=[])
```

### HTTP/2 Protocol

https://tkstorm.com/posts-list/programming/http/http2-review/