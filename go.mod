module github.com/https-whoyan/MafiaBotHelper

go 1.22

require (
	github.com/disgoorg/disgo v0.18.13
	github.com/disgoorg/snowflake/v2 v2.0.3
	github.com/https-whoyan/MafiaBot v0.0.1
	github.com/https-whoyan/MafiaCore v0.0.1
	github.com/joho/godotenv v1.5.1
	github.com/valyala/fasthttp v1.57.0
	golang.org/x/sync v0.8.0
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/disgoorg/json v1.2.0 // indirect
	github.com/fasthttp/router v1.5.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/sasha-s/go-csync v0.0.0-20240107134140-fcbab37b09ad // indirect
	github.com/savsgio/gotils v0.0.0-20240704082632-aef3928b8a38 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
)

replace github.com/https-whoyan/MafiaCore v0.0.1 => ../MafiaCore

replace github.com/https-whoyan/MafiaBot v0.0.1 => ../MafiaBot
