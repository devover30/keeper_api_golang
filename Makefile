build-prod:
	sudo docker build -t keeper_app .
run-db:
	sudo docker compose -f docker-compose.db.yml up --remove-orphans
stop-db:
	sudo docker compose -f docker-compose.db.yml down
build-tar:
	docker save -o keeper_app.tar keeper_app
