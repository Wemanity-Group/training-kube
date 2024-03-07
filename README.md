# Go

Vous devez d'abord créer un module Go :

go mod init go-backend

Un fichier go.mod est créé.

Pour activer CORS, on va avoir besoin du package github.com/gorilla/mux dans le Dockerfile :

go get github.com/gorilla/mux
go mod download github.com/gorilla/mux

docker build -t backend-app .

docker run -p 8080:8080 backend-app

http://localhost:8080/

404 page not found

http://localhost:8080/hello

Hello, Docker Multistage!

# ReactJS

Créer le fichier frontend-app/src/Hello.js :

```
// src/Hello.js
import React, { useState, useEffect } from 'react';

function Hello() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    async function fetchMessage() {
      const response = await fetch('/hello');
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

Dans le fichier App.js, modifier :

import Hello from './Hello';
return (
    <div className="App">
      <header className="App-header">
        <Hello />
      </header>
    </div>
  );

npm start


Si :
Hello from Frontend!
Response from Backend:

Le backend ne répond pas. Est-il lancé ?

docker run -p 8080:8080 backend-go

On obtient :

Hello from Frontend!
Response from Backend: Hello from Backend!

npm run build

docker build -t frontend-app .

docker run -p 8081:80 frontend-app

docker-compose up

```
Hello from Frontend!
Response from Backend: Hello from Backend!
```

----------





kubectl config use-context docker-desktop
helm install nginx-demo ./nginx-demo
