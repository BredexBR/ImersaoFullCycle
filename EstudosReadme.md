# Imersão FullStack && FullCycle
Refere-se a imersao da FullCycle utilizando as tecnologias: 
- Typescript/Javascript
- Nest.js
- Prisma
- API Rest
- Autenticação(Kong)

## Objetivo:
- Desenvolver um sistema de venda de ingresos com grande demanda de clientes utilizando o site simultaneamente.
- Gerenciamento de eventos/processamento de reservas.
- integração com sistemas de parceiros.


## Bibliotecas/Ferramentas
### Iniciando:
- Comando para instalar o Nest:
   ```bash
   npm install -g @nestjs/cli

- Comandos para gerar o projeto com configurações Nest:
   ```bash   
   nest new nestjs-partners-api

- Comando para instalar as configurações do nest(para variáveis de ambiente):
   ```bash
   npm install @nest.js/config

- Comando para usar ferramenta que permite carregar variáveis de ambiente a partir de um arquivo .env
   ```bash
   npm install dotenv-cli

- Para iniciar o projeto:
   ```bash
   npm install -g @nestjs/cli

- Após alterar pasta package.json relacionado aos partners rodar:
   ```bash
   npm run migrate:partner1
   npm run migrate:partner2

- Acesso em seu navegador para ver se esta rodando o sistema:
   ```bash 
   http://localhost:3000/

### Nest:
- Para gerar novos CRUDS basta inserir:
   ```bash
   nest g resource 
- escolha o nome do recurso e selecione a opção "REST API".

- Para gerar um modulo do prisma:
   ```bash
   nest g module prisma

- Para gerar o serviço para o prisma:
   ```bash
   nest g service prisma

- Para gerar uma biblioteca:
   ```bash
   nest g library

- Após isso escolha o nome "core" e pressione enter novamente(caso escreve alguma palavra obtera um prefixo)..

### Docker:
- Para gerar o SQL com o Docker basta adicionar as configurações em um arquivo .yaml como por exemplo docker-compose.yaml e inserir o comando:
   ```bash
   docker compose up

- Caso deseje excluir os container basta inserir:
   ```bash
   docker compose down

- Caso deseje conferir as databases do banco de dados basta inserir:
   ```bash
   docker compose exec db bash
   mysql -uroot -proot
   show databases;

- Caso deseje acessar o app criado no docker-compose e configurado no Dockerfile:
   ```bash
   docker compose exec app bash

- após o comando anterior:
   ```bash
   npm install @prisma/client

### Dev Containers:
- Após a instalação va na pesquisa do VS CODE no centro superior e digite:
   ```bash
   >dev containers: Open folder in Containers

- Depois selecione a pasta do projeto nestjs-partners-api. Vai recarregar e abrir novamente o VSCode, mas dessa vez vai estar dentro do container do Docker. Copie os arquivos devcontainer.js e docker-composer.yaml desse projeto para o seu. Assim as configurações do terminal zsh e outras estarão configuradas no container. E por ultimo na pesquisa do VS CODE no centro superior para as salvar as alterações digite:
   ```bash
   >dev containers: rebuild Containers

### MySQL:
- Mostra as databases:
   ```bash
   show databases;

- Muda para o banco de dados nest:
   ```bash
   use nest;

- Demonstra as informações da tabela Event:
   ```bash
   describe Event;

### Prisma:
- Para criar as configurações do prisma dentro do projeto: 
   ```bash
   npx prisma init

- Para gerar as migrações do Prisma após configurar schema.prisma e .env:
   ```bash
   npx prisma migrate dev

### Kong
- Link para utilizar o site do Kong [https://konghq.com/](https://konghq.com/)

- para utilizar o comando docker-compose sera necessário instalar a ferramenta:
   ```bash
   sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

- de as permissões a pasta:
   ```bash
   sudo chmod +x /usr/local/bin/docker-compose

- confira se instalou vendo a versão:
   ```bash
   docker-compose --version

- Para utilizar o banco de dados configurado em docker-compose.with-db.yaml
   ```bash
   docker-compose -f docker-compose.with-db.yaml up -d

### Go
- Para instalar todas as ferramentas e fazer as atualizações do Go para utilizar no VS Code:
  Dentro do vs code aperte Ctrl+Shif+P escreva go: install/update tools selecione a caixinha branca ao lado esquerdo de todos.

- Para testar se tudo funcionou corretamente escreva "go" no terminal pressione enter e veja se aparece na primeira linha "Go is a tool for managing Go source code." e abaixo todas as informações.

- Para importar os arquivos dentro do programa:
   ```bash
   go mod init https://github.com/BredexBR/ImersaoFullCycle/golang

- A pasta "internal" no projeto vai ser o local onde sera configurado as bibliotecas internas.

- Para baixar as dependências dos projetos por exemplo(uuid):
   ```bash
   go mod tidy 
   
### Extensões do VS Code:
- Prisma (Suporte para utilizar o prisma no VS Code)
- ESLint (verificar possíveis erros de javaScript)
- Prettier - Code formatter (Auxiliar nas "boas práticas" do código)
- Docker (Poder ter uma visualização melhor do Docker dentro do GitHub)
- Live Preview (Poder ver paginas web no VS Code)
- REST Client (Poder testar mais facilmente as requisições http)
- WSL (Poder utilizar ambiente linux no VS Code)
- Dev Containers (Integração do VS Code com o DOCKER)
- GO (para poder utilizar a linguagem go dentro do VS Code)