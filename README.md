# Translator Monorepo

This open-source monorepo combines two main services:
1. **Translator Service (Python)** – uses the M2M100 model to translate text.
2. **Go API** – provides an HTTP endpoint that calls the Translator service.

---

## Tech Specs
- **Python 3.11** for the Translator Service
- **Golang 1.22** for the Go API
- **Minimum Specs**: 2 CPU / 4 GB RAM
- **Docker** >= 20.x
- **Docker Compose** >= v2

---

## Container Images
Images are hosted on a Harbor-like registry:

- **Translator Service**: `https://registry.mirats.cloud/openkits/translator-service:1.0.0`
- **Go API**: `https://registry.mirats.cloud/openkits/go-api:1.0.0`

The `docker-compose.yml` references these images directly.

---

## Repository Structure

```
translator-monorepo
├── .gitignore
├── .dockerignore
├── README.md
├── docker-compose.yml
├── install.sh
├── build-images.sh       # Private script (excluded from Git)
├── translator
│   ├── Dockerfile
│   ├── requirements.txt
│   └── translator.py
└── api
    ├── Dockerfile
    ├── main.go
    ├── go.mod
    └── go.sum
```

### `.gitignore`
Excludes OS junk, compiled artifacts, logs, and `build-images.sh`.

### `.dockerignore`
Prevents Git, node_modules, and build artifacts from bloating the Docker image.

### `install.sh`
Automates Docker & Docker Compose installation if missing, then pulls and runs the services.

### `build-images.sh`
For maintainers to locally build & push images to the registry. Excluded from Git by default.

---

## Setup & Usage

1. **Clone the Repository**:
    ```bash
    git clone <repo-url>
    cd translator-monorepo
    ```

2. **Run the Installer**:
    ```bash
    chmod +x install.sh
    ./install.sh
    ```
    This checks/install Docker & Docker Compose if needed, then pulls and starts the containers in detached mode.

3. **Access the Services**:
    - **Go API**: [http://localhost:8080/api/v1/translate](http://localhost:8080/api/v1/translate)
    - **Translator** (Python FastAPI): [http://localhost:8000/translate](http://localhost:8000/translate)

---

## Testing the Translation Endpoint

Send a POST request to `go-api`:
```bash
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"text": "Hello, world!", "target_lang": "fr"}' \
     http://localhost:8080/api/v1/translate
```

**Sample Response**:
```json
{
  "translated_text": "Bonjour le monde!",
  "detected_source_lang": "en"
}
```

---

## Managing Containers

- Stop containers:
  ```bash
  docker-compose down
  ```
- Check logs:
  ```bash
  docker-compose logs -f
  ```

---

## Building & Pushing Images (Maintainers Only)

Use `build-images.sh` for internal workflows:

```bash
chmod +x build-images.sh
./build-images.sh
```

This will:
1. Build & push `translator-service`.
2. Build & push `go-api`.

Because `build-images.sh` is private, it’s excluded via `.gitignore`.

---

## Contributing

Feel free to open PRs or issues for bug fixes, improvements, or new features. We welcome community support!

---

## License
This project is [MIT Licensed](https://opensource.org/licenses/MIT).
