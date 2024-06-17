# Weather App by CEP

## Descrição

Este é um sistema desenvolvido em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O sistema é implantado no Google Cloud Run.

## Como Executar

### Requisitos

- Docker
- Docker Compose
- Conta no Google Cloud

### Passos para Executar Localmente

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/AnaGFranco/github.com/leocastr0/weather_by_cep.git
   cd github.com/leocastr0/weather_by_cep
   ```

2.**Construir as Imagens**: Na raiz do projeto, execute:

   ```
   docker-compose build
   ```

3. **Executar o Docker Compose**: Na raiz do projeto (onde está o arquivo `docker-compose.yml`), execute:

   ```
   docker-compose up
   ```


4. **Acesse a aplicação:**

   A aplicação estará disponível em `http://localhost:8080`.

### Exemplos de Requisição

- **Para obter o clima de um CEP específico:**

   ```sh
   curl -X GET http://localhost:8080/weather/29902555
   ```

### Testando o Serviço no Google Cloud Run

URL do Google Cloud Run. Use essa URL para fazer requisições:

```sh
curl -X GET  https://github.com/leocastr0/weather_by_cep-xhekudbara-uc.a.run.app/weather/29902555
```

```sh
curl -X GET https://github.com/leocastr0/weather_by_cep-xhekudbara-uc.a.run.app/weather/01153000
```

Código HTTP: 404
Mensagem: can not find zipcode
<img width="1373" alt="image" src="https://github.com/AnaGFranco/github.com/leocastr0/weather_by_cep/assets/55562874/134d9cf6-0870-4683-bc4b-c35e591dcc31">


Código HTTP: 422
Mensagem: invalid zipcode
<img width="1373" alt="image" src="https://github.com/AnaGFranco/github.com/leocastr0/weather_by_cep/assets/55562874/53bead93-1a0a-4949-bf56-4ec7f321c320">

Código HTTP: 200
Response Body: { "temp_C": 15, "temp_F": 59, "temp_K": 288.15 }
<img width="1373" alt="image" src="https://github.com/AnaGFranco/github.com/leocastr0/weather_by_cep/assets/55562874/cbcc5a1e-8c52-4107-9c9b-b25a23ff9a9c">


