ENV_GATE=CONF_PROVIDER_FILE_PATH=etc/gate/server.yml \
		CONF_CONSUMER_FILE_PATH=etc/gate/client.yml \
		APP_LOG_CONF_FILE=etc/gate/log.yml

ENV_GAME=CONF_PROVIDER_FILE_PATH=etc/game/server.yml \
		CONF_CONSUMER_FILE_PATH=etc/game/client.yml \
		APP_LOG_CONF_FILE=etc/game/log.yml

.PHONY: run-gate
run-gate:
	$(ENV_GATE) go run *.go -t gate

.PHONY: run-game
run-game:
	$(ENV_GAME) go run *.go -t game

.PHONY: run
run:
	make run-game run-gate -j2