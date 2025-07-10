# Frontend: Sistema de Interesse de Matrícula (UnB)

Esta aplicação é a interface de usuário (UI) para o Sistema de Interesse de Matrícula. Ela foi construída com Next.js e TypeScript para fornecer uma experiência de usuário moderna e reativa, consumindo a API de backend em Go.

## Tecnologias Utilizadas

- **Framework:** [Next.js](https://nextjs.org/) (com App Router)
- **Linguagem:** [TypeScript](https://www.typescriptlang.org/)
- **Estilização:** [Tailwind CSS](https://tailwindcss.com/)
- **Qualidade de Código:** [ESLint](https://eslint.org/)

## Como Executar (Ambiente de Desenvolvimento)

Para rodar este projeto de frontend, certifique-se de que o **backend já esteja em execução**, pois esta aplicação precisa da API para funcionar.

1.  **Navegue para a pasta do frontend** (a partir da raiz do monorepo):
    ```bash
    cd frontend
    ```

2.  **Instale as dependências** do projeto (só precisa fazer isso uma vez):
    ```bash
    npm install
    ```

3.  **Execute o servidor de desenvolvimento:**
    ```bash
    npm run dev
    ```

4.  Abra seu navegador e acesse [http://localhost:3000](http://localhost:3000) para ver a aplicação.

## Estrutura de Pastas (`src/`)

- **/app:** Contém as rotas e páginas da aplicação, seguindo o padrão do App Router.
- **/components:** Contém componentes React reutilizáveis (botões, formulários, etc.).
- **/services:** Contém a lógica para fazer chamadas à nossa API em Go.
- **/lib:** Contém funções de utilidade geral.