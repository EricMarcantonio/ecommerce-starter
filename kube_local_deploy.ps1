Set-Location deployment
kubectl.exe delete secrets --all
kubectl.exe delete config --all
kubectl.exe delete clusterrole --all
kubectl.exe delete deployments --all
kubectl.exe delete all --all
kubectl.exe apply -f 'configmap.yaml,secrets.yaml'
kubectl.exe apply -f .
#minikube.exe service server-external-service