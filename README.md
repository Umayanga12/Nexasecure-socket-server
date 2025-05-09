# 🚀 Nexasecure Socket Server

![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/github/license/Umayanga12/Nexasecure-socket-server)
![Build](https://img.shields.io/badge/build-passing-brightgreen)
![Contributions](https://img.shields.io/badge/contributions-welcome-ff69b4)

A **high-performance socket server** built with Go, designed for **secure** and **scalable** real-time communication.


## ✨ Features

- ⚡ **Real-time Communication** – Efficient and reliable data exchange between clients.
- 🧠 **Redis Integration** – Utilizes [go-redis](https://github.com/go-redis/redis) for caching and pub/sub messaging.
- 🪵 **Structured Logging** – Powered by [zap](https://github.com/uber-go/zap) for high-performance, leveled logs.
- ⚙️ **Environment Management** – Easy config setup via [godotenv](https://github.com/joho/godotenv).


## 📦 Installation

1. **Clone the repository**  
   ```bash
   git clone https://github.com/Umayanga12/Nexasecure-socket-server.git
   cd nexasecure-socket-server

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Build the server**

   ```bash
   go build -o socket-server
   ```


## 🧪 Usage

1. **Configure environment**
   Create a `.env` file with the following variables:

   ```env
   REDIS_HOST=localhost
   REDIS_PORT=6379
   LOG_LEVEL=info
   ```

2. **Start the server**

   ```bash
   ./socket-server
   ```

3. **Connect using a client**
   Use any TCP client to connect and exchange data securely and efficiently.

## 🛠 Development

### ✅ Prerequisites

* Go **1.24+**
* Redis server running locally or remotely

### 🧪 Run Tests

```bash
go test ./...
```

### 🧹 Format Code

```bash
go fmt ./...
```

## 🤝 Contributing

We ❤️ contributions!

1. Fork the repository.
2. Create a new branch:

   ```bash
   git checkout -b feature-name
   ```
3. Make your changes & commit:

   ```bash
   git commit -m "Add feature-name"
   ```
4. Push your branch:

   ```bash
   git push origin feature-name
   ```
5. Open a **Pull Request**.


## 🙏 Acknowledgments

* [Zap](https://github.com/uber-go/zap) – Fast and structured logging.
* [Go-Redis](https://github.com/go-redis/redis) – Redis client library.
* [Godotenv](https://github.com/joho/godotenv) – Environment configuration made easy.
