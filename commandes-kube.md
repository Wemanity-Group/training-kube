# Commandes pour jouer Kube

Passer en admin avec l'outil Wema;

Installer Brew :

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Installer kubectl :

```
brew install kubectl
```

Vérifier l’installation :

```
kubectl version --client
```

Installer Docker Desktop : `https://docs.docker.com/desktop/install/mac-install/`

## Configurer Kubernetes dans Docker Desktop

Télécharger Docker Desktop : https://www.docker.com/products/docker-desktop

Dans les préférences Docker Desktop, accédez à la section "Kubernetes" et activer Kubernetes.

Vérifier que Kubernetes est bien installé avec la commande :

```
kubectl get nodes
```

En cas de problème :

```
rm ~/.kube/config
kubectl config unset current-context
# Redémarrer Docker Desktop
kubectl config view
```

## Installation de Lens

Télécharger Lens depuis https://k8slens.dev/

## Création des objets

```
kubectl apply -f kubernetes/backend-deployment.yaml
```

Vérifier que les pods sont Running.

Une erreur survient : l'image n'est pas trouvée. La raison est que Kubernetes essaie de la télécharger depuis Internet, depuis Docker Hub. Il faut modifier l'attribut :

```
imagePullPolicy: Never
```

ou alors utiliser un tag spécifique qui n'existe pas sur Internet.

```
kubectl apply -f kubernetes/backend-service.yaml
```

```
kubectl apply -f kubernetes/frontend-deployment.yaml
kubectl apply -f kubernetes/frontend-service.yaml
```

Vérifier que les pods sont Running.

```
kubectl apply -f kubernetes/ingress.yaml
```

Ajouter le nom de domaine au fichier `/etc/hosts` :

```
sudo nano /etc/hosts
```

Ajouter au fichier :

```
127.0.0.1    your-domain.com
```

^o puis Entrée pour écrire la sortie.
^x pour sortir.


Commandes utiles :

```
kubectl get pods
kubectl get svc
kubectl exec -it <nom-du-pod> -- /bin/bash
apt update
apt install curl
```


Vérifier un service :

```
kubectl get svc frontend-service
curl http://10.111.115.156:80
curl http://<url-de-l-ingress>
kubectl logs <nom-du-pod>
```

Tout redéployer :

```
kubectl apply -f kubernetes/backend-deployment.yaml
kubectl apply -f kubernetes/backend-service.yaml
kubectl apply -f kubernetes/frontend-deployment.yaml
kubectl apply -f kubernetes/frontend-service.yaml
kubectl apply -f kubernetes/ingress.yaml
```

Entrer dans le backend :

Vérifier qu'on accède bien au frontend :

```
curl http://backend:8080
curl http://frontend:80
```

Par défaut, il n'y a pas d'ingress controller dans Docker Desktop. Il faut en installer un soit-même :

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.6.4/deploy/static/provider/cloud/deploy.yaml
```

Vérifier :

```
kubectl -n ingress-nginx get pod
```

Comme on utilise l'ingress Nginx, ajouter au fichier ingress.yaml :

```
  annotations:
    kubernetes.io/ingress.class: nginx
```

Puis vérifier que ça fonctionne :

```
curl your-domain.com
```

Modifier la route appellée dans Hello.js.
