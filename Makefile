build-dev:
	docker-compose --verbose -f docker-compose-dev.yml -p fedha build

run-dev:
	docker-compose --verbose  -f docker-compose-dev.yml -p fedha up --build --remove-orphans

prune-dev:
	docker-compose -f docker-compose-dev.yml -p fedha down --remove-orphans --volumes --rmi all

