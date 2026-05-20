# gin-gomic

Un projet d'apprentissage Go construit avec **Gin** et **GORM** — une API REST avec authentification JWT, accès base de données MySQL et déploiement Docker.

> 🎓 Réalisé dans le cadre de mon apprentissage de Go, Gin et GORM.

---

## Stack technique

| Couche | Technologie |
|---|---|
| Langage | Go 1.25 |
| Framework HTTP | [Gin](https://github.com/gin-gonic/gin) v1.11 |
| ORM | [GORM](https://gorm.io) v1.31 + driver MySQL |
| Base de données | MySQL 8.0 |
| Authentification | JWT ([golang-jwt/jwt](https://github.com/golang-jwt/jwt) v5) |
| Configuration | `.env` via [godotenv](https://github.com/joho/godotenv) |
| Conteneurisation | Docker + Docker Compose |

---

## Structure du projet

```
gin-gomic/
├── cmd/
│   └── api/
│       └── main.go       # Point d'entrée de l'application
├── config/               # Chargement de la configuration
├── internal/             # Logique métier (handlers, services, modèles)
├── pkg/                  # Packages partagés / utilitaires
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── .env                  # Variables d'environnement (non versionné)
```

---

## Prérequis

- [Go](https://go.dev/dl/) >= 1.21
- [Docker](https://www.docker.com/) et Docker Compose

---

## Démarrage rapide

### Avec Docker Compose (recommandé)

Lance l'API, la base de données MySQL et Adminer en une seule commande :

```bash
docker compose up --build
```

| Service | URL |
|---|---|
| API | http://localhost:8080 |
| Adminer (UI MySQL) | http://localhost:7080 |
| MySQL | localhost:3312 |

### En local (sans Docker)

1. Copie le fichier d'exemple et renseigne tes variables :
   ```bash
   cp .env.example .env
   ```

2. Installe les dépendances :
   ```bash
   go mod download
   ```

3. Lance l'application :
   ```bash
   go run ./cmd/api/main.go
   ```

---

## Variables d'environnement

Crée un fichier `.env` à la racine du projet :

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=test

JWT_SECRET=your_secret_key

APP_PORT=8080
```

---

## Ce que j'ai appris

- Structurer une API REST en Go avec le pattern `cmd/internal/pkg`
- Utiliser **Gin** pour le routing, les middlewares et la validation des requêtes
- Manipuler une base de données MySQL avec **GORM** (migrations, CRUD)
- Implémenter une authentification par **JWT**
- Containeriser une application Go avec un **Dockerfile multi-stage**
- Orchestrer les services avec **Docker Compose**

---

## Auteur

**cabitibaly** — [GitHub](https://github.com/cabitibaly)