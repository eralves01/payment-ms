# Payment Microservice

Este repositório contém um microserviço de pagamentos desenvolvido em Go, utilizando PostgreSQL como banco de dados e PgAdmin como ferramenta de administração. O ambiente de desenvolvimento é gerenciado com Docker e Docker Compose, incluindo suporte para hot reload com a ferramenta `air`.

## **Requisitos**
Antes de iniciar, certifique-se de ter instalado em sua máquina:
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

## **Serviços**
O projeto utiliza os seguintes serviços Docker:

1. **PostgreSQL (`postgres`)**
   - Banco de dados utilizado pelo microserviço de pagamentos.
   - Executa na porta `5432`.
   - Os dados persistem no volume `postgres_data`.

2. **PgAdmin (`pgadmin`)**
   - Interface gráfica para gerenciar o banco de dados PostgreSQL.
   - Acessível via navegador em `http://localhost:5050`.
   - Credenciais padrão:
     - **E-mail:** `admin@admin.com`
     - **Senha:** `123456`

3. **Aplicação Go (`app`)**
   - Implementa a lógica do microserviço de pagamentos.
   - Utiliza `air` para hot reload.
   - Executa na porta `8080`.

## **Como Inicializar o Sistema**
Siga as etapas abaixo para configurar e rodar o sistema localmente.

### **1. Clonar o repositório**
```sh
git clone https://github.com/seu-usuario/payment-ms.git
cd payment-ms
```

### **2. Criar e configurar o `.env`**
Crie um arquivo `.env` na raiz do projeto com as credenciais do banco de dados:
```
POSTGRES_USER=eralves
POSTGRES_PASSWORD=123456
POSTGRES_DB=payment_db
PGADMIN_DEFAULT_EMAIL=eralves01@gmail.com
PGADMIN_DEFAULT_PASSWORD=123456
```

### **3. Subir os containers**
```sh
docker-compose up --build
```
O `--build` é usado para garantir que a imagem da aplicação seja reconstruída corretamente.

### **4. Acessar os serviços**
- **Banco de Dados:**
  - Host: `localhost`
  - Porta: `5432`
  - Usuário: `eralves`
  - Senha: `123456`
  - Banco: `payment_db`

- **PgAdmin:**
  - URL: [http://localhost:5050](http://localhost:5050)
  - E-mail: `eralves01@gmail.com`
  - Senha: `123456`

- **API Go:**
  - Rodando em `http://localhost:8080`

### **5. Parar os containers**
Para parar os serviços, pressione `CTRL+C` ou rode:
```sh
docker-compose down
```

### **6. Rodar em background (modo daemon)**
Se quiser rodar os containers em segundo plano, use:
```sh
docker-compose up -d
```

## **Hot Reload com Air**
A aplicação Go usa o `air` para recarregar automaticamente quando o código é modificado. Se quiser rodar a aplicação fora do Docker, instale o `air` localmente:
```sh
go install github.com/air-verse/air@latest
```
Depois, inicie o servidor com:
```sh
air
```

## **Dicas**
- Se houver erro ao rodar os containers, tente remover os volumes:
  ```sh
  docker-compose down -v
  ```
- Para visualizar logs dos containers:
  ```sh
  docker-compose logs -f
  ```

Agora você já pode rodar o sistema localmente para desenvolvimento! 🚀

 
