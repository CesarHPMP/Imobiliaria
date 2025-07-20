# 🏠 Imobiliária CRM

**Imobiliária CRM** é uma aplicação web desenvolvida em Go com foco no gerenciamento de usuários e propriedades para imobiliárias. A aplicação oferece autenticação via JWT, endpoints RESTful seguros e integração com banco de dados PostgreSQL.

![Go](https://img.shields.io/badge/Go-1.20-blue?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15-blue?logo=postgresql)
![Docker](https://img.shields.io/badge/Docker-%231572B6.svg?logo=docker&logoColor=white)
![License](https://img.shields.io/github/license/CesarHPMP/Imobiliaria)

---

## ✨ Funcionalidades

- ✅ Cadastro e login de usuários com autenticação via JWT
- 🏡 Cadastro e listagem de propriedades
- 🔐 Proteção de rotas com middleware
- 📦 Conexão com banco de dados PostgreSQL
- 🐳 Docker + Docker Compose para ambiente de desenvolvimento
- 📁 Arquitetura modular e escalável em Go

---

## 📦 Estrutura do Projeto

```

Imobiliaria/
├── backend/
│   ├── main.go
│   ├── go.mod / go.sum
│   ├── internal/
│   │   ├── config/         # Configurações carregadas por struct
│   │   ├── controllers/    # Lógica dos endpoints
│   │   ├── database/       # Conexão e migrações
│   │   ├── middleware/     # Middleware de autenticação
│   │   ├── routes/         # Rotas públicas e protegidas
│   │   └── utils/          # Hashing, JWT, utilidades
├── docker/
│   └── docker-compose.yml
└── README.md

````

---

## 🚀 Como Rodar Localmente

### 🔧 Pré-requisitos

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- (Opcional) Go instalado localmente para desenvolvimento

### 📥 Clone o repositório

```bash
git clone https://github.com/CesarHPMP/Imobiliaria.git
cd Imobiliaria
````

### 🐳 Rode com Docker Compose

```bash
docker-compose -f docker/docker-compose.yml up --build
```

O backend estará disponível em `http://localhost:8080`.

---

## 🔐 Autenticação e Segurança

* Usuários são autenticados com **JWT**.
* Rotas que exigem autenticação são prefixadas por `/api/protected`.
* Exemplo de header necessário:

```http
Authorization: Bearer <seu_token_aqui>
```

---

## 🛠 Endpoints da API

| Método | Rota                            | Protegida | Descrição                 |
| ------ | ------------------------------- | --------- | ------------------------- |
| POST   | `/api/addUsers`                 | ❌         | Criação de novo usuário   |
| POST   | `/api/login`                    | ❌         | Login e geração de token  |
| GET    | `/api/protected/users`          | ✅         | Lista todos os usuários   |
| GET    | `/api/protected/properties`     | ✅         | Lista propriedades        |
| POST   | `/api/protected/createProperty` | ✅         | Cadastra nova propriedade |

---

## 🧪 Testes

*Testes automatizados ainda não incluídos.*
Você pode rodar testes manuais via [Insomnia](https://insomnia.rest/) ou [Postman](https://www.postman.com/).

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](./LICENSE) para mais detalhes.

---

## 👨‍💻 Autor

Desenvolvido por **César Henrique Policarpo de Melo**
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?logo=linkedin)](https://www.linkedin.com/in/cesarhpmp)

---

## 🙌 Contribuições

Sinta-se à vontade para abrir **Issues**, propor melhorias ou fazer um **Fork** do projeto.

---

> *"Código limpo é aquele que funciona — mas também é legível, seguro e preparado para crescer."* 🚀

