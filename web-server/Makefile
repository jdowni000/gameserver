build:
	docker build -f Dockerfile -t jdowni000/web-server:v1.0.0 .

push:
	docker push jdowni000/web-server:v1.0.0

test:
	go test ./...

run:
	docker run -p 8080:8080 jdowni000/web-server:v1.0.0

kubernetes:
	kubectl create deployment jdowni000-deployment --image=jdowni000/web-service:v1.0.0 --replicas=3 --port=8080 -oyaml --dry-run=client > kubernetes.yaml
