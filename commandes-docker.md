# Commandes à jouer pour Docker et Docker Compose

Les prérequis :
- avoir installé Go,
- avoir installé NodeJS,
- avoir installé ReactJS.

## Création de l'image Docker Backend

Setup :

```
mkdir backend-app/
cd backend-app/
```

Création d’un module Go :

```
go mod init go-backend
```

Le fichier go.mod est créé.

Pour activer CORS, on va avoir besoin du package github.com/gorilla/mux :

```
go get github.com/gorilla/mux
```

Le fichier go.sum est créé et le fichier go.mod est modifié. Ils référencent le module github.com/gorilla/mux.

```
go mod download github.com/gorilla/mux
```

Créer le fichier main.go avec le contenu fourni :

```
package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"message": "Hello from Backend!"}`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", helloHandler)

	// Ajoutez les options du middleware CORS ici
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	http.Handle("/", corsHandler(r))

	fmt.Println("Server is listening on port 80")
	http.ListenAndServe(":80", nil)
}
```

Créer le fichier Dockerfile :

```
FROM golang:1.16 AS build

WORKDIR /app

COPY main.go go.mod /app/

RUN go mod download github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

RUN go build -o hello-app

FROM debian:buster-slim

COPY --from=build /app/hello-app /usr/local/bin/

EXPOSE 80

CMD ["hello-app"]
```

Build de l’image Docker :

```
docker build -t backend .
```

L'image backend est crée :

```
docker images
```

Exécution du conteneur Docker sur le port 80 :

```
docker run -d -p 80:80 --name backend backend
```

Vérifier que le conteneur tourne :

```
docker ps -a | grep backend
```

Vérifier qu'il répond :

```
curl localhost
```

Il y a une erreur mais le serveur web répond :

```
404 page not found
```

Ce n'est pas la bonne route.

Vérifier qu'il répond sur la route `/hello` :

```
curl localhost/hello
```

Ou tester via un navigateur :

```
open http://localhost/hello
```

## Création de l'image Docker Frontend

Setup :

```
mkdir frontend-app/
cd frontend-app/
```

Créer un projet ReactJS :

```
npx create-react-app frontend-app
```

Une arborescence complète a été créée.

Vérifier que tout fonctionne correctement :

```
cd frontend-app
npm start
```

Le navigateur doit être lancé et la page ReactJS par défaut doit être affichée.

Dans le répertoire src, créer le fichier Hello.js :

```
import React, { useState, useEffect } from 'react';

function Hello() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    async function fetchMessage() {
      const response = await fetch('http://localhost/hello');
      const data = await response.json();
      setMessage(data.message);
    }

    fetchMessage();
  }, []);

  return (
    <div>
      <h1>Hello from Frontend!</h1>
      <p>Response from Backend: {message}</p>
    </div>
  );
}

export default Hello;
```

Modifier le fichier App.js :

- ajouter en début de fichier :

```
import Hello from './Hello';
```

- modifier le corps de la fonction App() :

```
return (
    <div className="App">
      <header className="App-header">
        <Hello />
      </header>
    </div>
  );
```

La page affichée par le navigateur web doit avoir changé et afficher "Hello from Frontend!".

Quelques explications sur ce qu'on vient de faire :
- Le fichier App.js est le point d'entrée principal de l'application ReactJS.
- `import Hello from './Hello';` : importe le composant Hello depuis le fichier Hello.js
- `function App() { ... }` : déclare la fonction `App`, qui est un composant fonctionnel React. Ce composant retourne un élément JSX qui représente la structure de l'interface utilisateur de l'application.
- `<Hello />` : L'élément Hello représente l'utilisation du composant Hello importé précédemment. Cela signifie que le composant Hello est rendu à l'intérieur de l'en-tête de l'application.

Créer le fichier `Dockerfile` avec ce contenu :

```
FROM nginx:alpine

RUN rm -rf /etc/nginx/conf.d

COPY nginx.conf /etc/nginx/conf.d/default.conf

COPY build /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
```

Voici quelques explications :
- `FROM nginx:alpine` : indique que nous voulons créer notre image à partir de l'image Docker officielle de Nginx basée sur Alpine Linux. Alpine Linux est une distribution Linux légère, ce qui permettra à notre image d'être petite et efficace.
- `RUN rm -rf /etc/nginx/conf.d` : supprime le répertoire /etc/nginx/conf.d dans l'image Docker. Par défaut, cet emplacement contient des fichiers de configuration Nginx. Nous allons le supprimer car nous voulons utiliser notre propre fichier de configuration.
- `COPY nginx.conf /etc/nginx/conf.d/default.conf` : copie le fichier de configuration Nginx nginx.conf dans le répertoire /etc/nginx/conf.d de l'image Docker. Ce fichier configure le serveur pour servir notre application web statique.
- `COPY build /usr/share/nginx/html` : copie le contenu de l'application web statique, qui est généré dans un répertoire build, dans le répertoire /usr/share/nginx/html de l'image Docker. Ca met donc les fichiers de l'application web à la disposition de Nginx pour les servir aux clients.
- `EXPOSE 80` : expose le port 80 sur le conteneur Docker. Cela signifie que lorsque le conteneur est lancé, le port 80 est accessible aux autres conteneurs ou à l'hôte sur lequel Docker est exécuté. Attention, cette instruction n'a qu'un impact de documentation, il n'a aucun autre impact.
- `CMD ["nginx", "-g", "daemon off;"]` : spécifie la commande par défaut à exécuter lorsque le conteneur est lancé. Dans ce cas, nous démarrons le serveur Nginx avec l'option -g daemon off; pour exécuter Nginx en mode non-démon. Cela permet à Docker de gérer correctement le processus et de le maintenir en cours d'exécution tant que le conteneur est actif.

On copie le fichier `nginx.conf`. Nous devons donc le créer :

```
server {
    listen 80;
    server_name localhost;

    location / {
        root /usr/share/nginx/html;
        index index.html;
        try_files $uri $uri/ /index.html;
    }
}
```

Voici des explications sur le fichier de configuration Nginx :
- Il doit se trouver dans le répertoire /usr/share/nginx/html.
- `server { ... }` : indique le début de la configuration d'un bloc de serveur dans Nginx.
- `listen 80;` : spécifie que le serveur Nginx écoute les requêtes entrantes sur le port 80.
- `server_name localhost;` : spécifie le nom du serveur auquel cette configuration s'applique. Dans ce cas, les requêtes pour le nom de domaine "localhost" sont gérées par ce serveur.
- `location / { ... }` : définit la configuration pour le traitement des requêtes entrantes pour l'URL racine ("/") du serveur.
- `root /usr/share/nginx/html;` : spécifie le répertoire racine à partir duquel les fichiers statiques sont servis (/usr/share/nginx/html ici).
- `index index.html;` : spécifie le fichier par défaut qui servi si l'URL demandée ne contient pas de nom de fichier spécifique.
- `try_files $uri $uri/ /index.html;` : définit la stratégie de traitement des requêtes qui ne correspondent pas à un fichier spécifique. Cela indique à Nginx de d'abord essayer de servir le fichier demandé ($uri), puis de regarder si une réécriture de l'URL existe ($uri/), et enfin de servir le fichier index.html s'il ne trouve aucune correspondance.

On va maintenant builder l'application ReactJS :

```
npm run build
```

Le répertoire `build` est créé et ne contient que des fichiers statiques.

Builder l'image Docker du Frontend à partir du répertoire parent :

```
docker build -t frontend .
```

Vérifier que le conteneur Docker fonctionne correctement :

```
docker run -p 8081:80 --name frontend frontend
```

Dans le navigateur web, aller à l'adresse http://localhost:8081/. Ceci doit être affiché :

```
Hello from Frontend!
Response from Backend: Hello from Backend!
```

## Arrêter les conteneurs

Arrêter l'application qui tourne localement sur le poste (^c).

Lister les conteneurs :

```
docker ps -a
```

Deux conteneurs doivent tourner : backend et frontend.

Supprimer les conteneurs :

```
docker rm frontend
docker stop backend
docker rm backend
```

Vérifier que tout est à l'arrêt :

```
docker ps -a
```

Plus aucun conteneur ne doit tourner.

Supprimer les images :

```
docker images
docker rmi frontend
docker rmi backend
```

Vérifier que les images ont bien été supprimées :

```
docker images
```

## Docker Compose

A la racine du projet, créer le fichier `docker-compose.yml` :

```
version: '3'
services:
  backend:
    container_name: backend
    build:
      context: ./backend-app
    image: backend
    ports:
      - "80:80"
  frontend:
    container_name: frontend
    build:
      context: ./frontend-app
    image: frontend
    ports:
      - "3000:80"
```

Créer les conteneurs :

```
docker-compose up -d
```

Comme les images Docker n'existent pas, Docker Compose va les créer automatiquement :

```
docker images
```

Puis vérifier que les conteneurs ont été créés :

```
docker ps -a
```

Dans le navigateur web, on a le résultat :

```
Hello from Frontend!
Response from Backend: Hello from Backend!
```

Pour déboguer, vous pouvez entrer dans le conteneur :

```
docker exec -it frontend sh
```









#kubectl config use-context docker-desktop
#helm install nginx-demo ./nginx-demo
