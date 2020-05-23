
dev:
	docker build -t discord-resource:dev ./

publish:
	docker build -t trivigy/discord-resource:latest ./ \
		&& docker push trivigy/discord-resource:latest
