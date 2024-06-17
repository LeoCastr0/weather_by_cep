# Weather App by CEP

## Descrição

Este é um sistema criado em Go que, dado um CEP, identifica a cidade correspondente e retorna o clima atual (temperatura em Celsius, Fahrenheit e Kelvin). A aplicação está implantada no Google Cloud Run.

## Como Executar

### Requisitos

- Docker
- Docker Compose
- Conta no Google Cloud

### Passos para Executar Localmente

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/leocastr0/weather_by_cep.git
   cd github.com/leocastr0/weather_by_cep
   ```

2.**Construir as Imagens**: Na raiz do projeto, execute:

   ```
   docker-compose build
   ```

3. **Executar o Docker Compose**: Na raiz do projeto (onde está o arquivo docker-compose.yml), execute:

   ```
   docker-compose up
   ```


4. **Acesse a aplicação:**

   A aplicação estará disponível em `http://localhost:8080`.
