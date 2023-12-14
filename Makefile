# 起動中のapi_serverのコンテナに接続する
bash_api_server:
	docker exec -it pinnacle-play-api_server /bin/bash

# 起動中のweb_frontのコンテナに接続する
bash_web_front:
	docker exec -it pinnacle-play-web_front /bin/bash