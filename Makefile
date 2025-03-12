.PHONY: setup run stop logs

setup:
	@echo "Checking and installing Docker Compose..."
	@if ! command -v docker-compose &> /dev/null; then \\
		curl -SL https://github.com/docker/compose/releases/latest/download/docker-compose-$$(uname -s)-$$(uname -m) -o /usr/local/bin/docker-compose && \\
		chmod +x /usr/local/bin/docker-compose; \
	else \
		echo \"Docker Compose already installed.\"; \
	fi

run: setup
	docker-compose up -d

stop:
	docker-compose down

logs:
	docker-compose logs -f
