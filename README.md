# Projeto de Banco de Dados: Sistema de Interesse de Matrícula (UnB)

Este projeto é uma aplicação full-stack desenvolvida como parte da disciplina de Banco de Dados. O objetivo é criar um sistema que permita aos alunos da Universidade de Brasília (UnB) manifestarem interesse prévio em disciplinas, ajudando a gestão acadêmica a prever a demanda e os alunos a planejarem melhor seus semestres.

Este repositório é um **monorepo**, contendo dois projetos principais:
- **/backend:** Uma API RESTful escrita em Go.
- **/frontend:** Uma aplicação web escrita em Next.js (React).

## Problema Resolvido

O sistema busca mitigar a ansiedade e a incerteza do período de matrícula, fornecendo uma plataforma onde a demanda por turmas pode ser visualizada antes do início do processo oficial. Isso permite que alunos, professores e coordenadores tomem decisões mais informadas.

## Tecnologias Utilizadas

#### Backend
- **Linguagem:** Go (Golang)
- **Banco de Dados:** PostgreSQL
- **Roteador HTTP:** Chi
- **Driver PostgreSQL:** pgx
- **Documentação da API:** Swagger (OpenAPI)

#### Frontend
- **Framework:** Next.js (React)
- **Linguagem:** TypeScript
- **Estilização:** Tailwind CSS

## Como Executar o Projeto Completo (Setup)

Siga os passos abaixo para configurar e executar o ambiente de desenvolvimento completo.

### 1. Pré-requisitos
* [Go](https://go.dev/doc/install) (versão 1.18+)
* [PostgreSQL](https://www.postgresql.org/download/) (versão 14+)
* [Node.js](https://nodejs.org/) (versão 18+)
* [Git](https://git-scm.com/downloads)

### 2. Clone o Repositório
```bash
git clone [https://github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB.git](https://github.com/JoaoJorgedosAnjos/sistema-de-interesse-de-matricula-UnB.git)
cd sistema-de-interesse-de-matricula-UnB
```

### 3. Configure o Banco de Dados
1.  Certifique-se de que seu servidor PostgreSQL está rodando.
2.  Use o "Script Mestre" (`master_script.sql` que preparamos) para criar e popular o banco de dados `unb_database`.
3.  No diretório `/backend`, renomeie `.env.example` para `.env` e preencha com suas credenciais do banco de dados (se estiver a usar variáveis de ambiente).

### 4. Execute a Aplicação
Você precisará de **dois terminais** abertos, ambos na raiz do projeto (`/sistema-de-interesse-de-matricula-UnB`).

**No Terminal 1 - Inicie o Backend:**
```bash
# Navegue para a pasta do backend
cd backend

# Baixe as dependências do Go
go mod tidy

# Inicie o servidor da API
go run ./cmd/api/main.go
```
> O backend estará rodando em `http://localhost:8080`. Você verá a documentação interativa em `http://localhost:8080/swagger/index.html`.

**No Terminal 2 - Inicie o Frontend:**
```bash
# Navegue para a pasta do frontend
cd frontend

# Instale as dependências do Node.js
npm install

# Inicie o servidor de desenvolvimento
npm run dev
```
> O frontend estará rodando em `http://localhost:3000`.

### 5. Acesse a Aplicação
Abra seu navegador e acesse **[http://localhost:3000](http://localhost:3000)** para usar a interface gráfica do sistema.