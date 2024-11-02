**prerequisites :upside_down_face: **

1. Docker
2. Helm
3. Minikube

**Steps to get server up :rocket: **
1. clone this repo
2. first build the docker image
```
docker build . -t order_service-app
```
3. Install the charts
```
helm install database ./database
helm install api ./api
```
4. expose the port of api server 
```
kubectl port-forward deployment/api 5000 5000
```
5. verify the changes on the postman :blush:
