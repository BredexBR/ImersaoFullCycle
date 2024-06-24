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
- Integração com sistemas de parceiros.

## Bibliotecas/Ferramentas
### Docker
- Para utilizar a configuração dentro do docker-compose.yaml de host, editar o arquivo no terminal(como administrador)
   ```bash
   cd C:\Windows\System32\drivers\etc>

- Após entrar nesse diretório
   ```bash
   notepad hosts

- E adicionar a linha:
   ```bash
   127.0.0.1 host.docker.internal

- Na pasta raiz remova todos os containers criados no docker(ativos e inativos):
   ```bash
   docker rm $(docker ps -a -q)

- Para Conferir se realmente nã possui nenhum container aberto:
   ```bash
   docker ps

- E crie o container para a aplicação:
   ```bash
   docker compose up

- Deixe esse terminal aberto(caso queira, deixe o nome salvo como "docker")

### Go
- Para instalar todas as ferramentas e fazer as atualizações do Go para utilizar no VS Code:
  Dentro do vs code aperte Ctrl+Shif+P escreva go: install/update tools selecione a caixinha branca ao lado esquerdo de todos.

- Para testar se tudo funcionou corretamente escreva "go" no terminal pressione enter e veja se aparece na primeira linha "Go is a tool for managing Go source code." e abaixo todas as informações.

- Para rodar a aplicação Go dentro da pasta golang digite no terminal:
   ```bash
   go run cmd/events/main.go

- Deixe este terminal aberto(caso queira, deixe o nome salvo como "golang")

### Nest:
- Comando para instalar o Nest:
   ```bash
   npm install -g @nestjs/cli

- Comando para usar ferramenta que permite carregar variáveis de ambiente a partir de um arquivo .env
   ```bash
   npm install dotenv-cli

- Para rodar a aplicação nest dentro de sua pasta "nestjs-partners-api" digite no terminal:
   ```bash
   docker compose exec nestjs bash

- Após isso:
   ```bash
   npm run migrate:partner1
   npm run migrate:partner2

- Para criar os dados do projeto(partner1):
   ```bash
   npm run start partner1-fixture

- E também(partner2):
   ```bash
   npm run start partner2-fixture

- Para rodar o WebServer (localhost:3000/events):
   ```bash
   npm run start:dev

- E deixe esse terminal aberto(caso queira, deixe o nome salvo como "partner1")

- Abra outro terminal e agora faça para o partner2 (localhost:3001/eventos):
   ```bash
   npm run start:dev partner2

- E também deixe esse terminal aberto(caso queira, deixe o nome salvo como "partner2")

### Next
- Para rodar a aplicação Next:
   ```bash
   docker compose exec nextjs bash
   npm run dev

- Deixe esse terminal aberto(caso queira, deixe o nome salvo como "nextjs")

### Kong
- O arquivo localizado em "kong-api-gateway/kong/declarative/kong.yaml" esta configurado todos os requisitos necessários para rodar a aplicação. Caso queira ver acesse no navegador:
   ```bash
   localhost:8002

   

### Extensões utilizadas do VS Code:
- Prisma (Suporte para utilizar o prisma no VS Code)
- ESLint (verificar possíveis erros de javaScript)
- Prettier - Code formatter (Auxiliar nas "boas práticas" do código)
- Docker (Poder ter uma visualização melhor do Docker dentro do GitHub)
- Live Preview (Poder ver paginas web no VS Code)
- REST Client (Poder testar mais facilmente as requisições http)
- WSL (Poder utilizar ambiente linux no VS Code)
- Dev Containers (Integração do VS Code com o DOCKER)
- GO (para poder utilizar a linguagem go dentro do VS Code)
- Tailwind CSS IntelliSense (trabalha com classes utilitárias(escondem propriedades de CSS))