# https://github.com/ogen-go/ogen
gen-by-oapi-codegen:
	docker compose run develop-service \
		oapi-codegen \
			--config=./config/oapi-codegen.yaml \
			./oas.yaml
air:
	docker compose run --publish 80:80 develop-service \
		air \
			-c ./config/.air.toml 
