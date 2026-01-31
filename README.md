# Chaos API 🚀

[![Go Reference](https://pkg.go.dev/badge/chaos-api.svg)](https://pkg.go.dev/chaos-api)  
[![Go Report Card](https://goreportcard.com/badge/github.com/seu-usuario/chaos-api)](https://goreportcard.com/report/github.com/seu-usuario/chaos-api)

Simula falhas e latência de serviços externos para testar aplicações de forma realista. Ideal para desenvolvimento de backend com **Go**, **chi** e **middleware de caos**.

---

## 🧠 Sobre

A **Chaos API** permite:

- Introduzir **latência** em requisições HTTP.
- Injetar **falhas aleatórias** (erros 500) em endpoints.
- Testar **timeout, retry e resiliência** em sistemas distribuídos.
- Configurar regras de caos dinamicamente via REST (`/chaos`).

Perfeito para **portfólio júnior Go**, mostrando backend de produção e middleware avançado.

---

## 📦 Estrutura do Projeto

```txt
chaos-api/
├── cmd/api/main.go            # Inicializa servidor, middleware e endpoints
├── internal/
│   ├── chaos/                # Lógica de caos (engine, configs, storage)
│   ├── handlers/             # Handlers HTTP (/users e /chaos)
│   ├── middleware/           # Middleware Chaos
│   └── server/               # Router chi + middlewares globais (CORS, logger)
├── go.mod
└── README.md

```
⚙️ Setup Rápido
Pré-requisitos

- Go 1.21+

- Git


Rodando localmente

```txt
git clone https://github.com/seu-usuario/chaos-api.git
cd chaos-api
go mod tidy
go run cmd/api/main.go
```
Servidor rodando em: http://localhost:8080
**ATENÇÃO:**
as regras de caos apenas começão a funcionas após sua cricação com post

🖥️ Endpoints

/users - Exemplo de endpoint com caos
```txt
GET /users
```

/chaos - CRUD de regras
Listar regras
```txt
GET /chaos
```
Criar regra
```txt
POST /chaos
Content-Type: application/json

{
  "path": "/users",
  "method": "GET",
  "error_rate": 0.3,
  "min_delay_ms": 200,
  "max_delay_ms": 1500
}
```

Deletar regra
```txt
DELETE /chaos/{id}
```

**Fluxo do Middleware**

1. Requisição chega no router chi

2. CORS processa primeiro (preflight OPTIONS seguro)

3. Chaos Middleware aplica:

- atraso

- erro 500

- ou passa normalmente

4. Handler responde