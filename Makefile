.PHONY: up down db_seed wait_health

up:
	docker compose up -d
	make wait_health
	make db_seed

down:
	docker compose down

wait_health:
	@echo "Waiting for Pinot to be ready..."
	@until [ $$(curl -s -o /dev/null -w "%{http_code}" http://localhost:9000/health) -eq 200 ]; do \
		echo "Pinot is not ready yet, waiting..."; \
		sleep 5; \
	done
	@echo "Pinot is ready."

db_seed:
	docker exec -it pinot-controller bin/pinot-admin.sh AddTable -tableConfigFile /config/table.json -schemaFile /config/schema.json -exec
	docker exec -it pinot-controller bin/pinot-admin.sh LaunchDataIngestionJob -jobSpecFile /config/job-spec.yml
