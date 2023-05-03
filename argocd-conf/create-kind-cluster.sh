kind create cluster --config ./clusterconfig-1.24.1.yaml --name argocd
kubectl create ns argocd
kubectl apply -f  -n argocd
echo "####### CREARTE NAMESPACE #############"
kubectl create ns org-147
kubectl create ns org-148
kubectl create ns argo-rollouts
kubectl create ns paas-application-staging
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
kubectl delete -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml
kubectl delete -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
kubectl delete serviceaccount -n kubernetes-dashboard admin-user
kubectl delete clusterrolebinding -n kubernetes-dashboard admin-user --clusterrole cluster-admin --serviceaccount=kubernetes-dashboard:admin-user
#kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
#token=$(kubectl -n kubernetes-dashboard create token admin-user)

