up:
	docker-compose up -d --build

down:
	docker-compose down

run:
	docker-compose up --build

install:
	cd front && npm install

clean-front:
	cd front && \
	rm yarn.lock && \
	rm -rf node_modules && \
	yarn install