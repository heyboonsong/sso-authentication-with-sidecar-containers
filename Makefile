
run_keycloak:
	kubectl apply -f ./k8s/service.yaml
	docker-compose -f ./keycloak/docker-compose.yml up -d

run_services:
	kubectl apply -f ./services/apps/album/k8s
	kubectl apply -f ./services/apps/comment/k8s
	kubectl apply -f ./services/apps/photo/k8s
	kubectl apply -f ./services/apps/post/k8s
	kubectl apply -f ./services/apps/todo/k8s
	kubectl apply -f ./services/apps/user/k8s

run_bff:
	kubectl apply -f ./bff/mobile-bff/k8s
	kubectl apply -f ./bff/web-bff/k8s

run_ingress:
	kubectl apply -f ./k8s/ingress.yaml
