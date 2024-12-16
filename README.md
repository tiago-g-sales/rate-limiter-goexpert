# weather-otel-goexpert
Sistema distribuido em 2 serviços que retornam o clima atual baseado em um CEP informado

# Desafio GOLang Observabilidade com trace distribuído - Consulta Temperatura baseado em um CEP informado - FullCycle 

Aplicação em Go sendo: 
  - Servidor HTTP Rest Client
  - Servidor HTTP Rest Server
  - Servidor Zipkin para apresentação do trace distribuído
  - Servidor Jaeger para apresentação do trace distribuído
  - Servidor Prometheus
  - Servidor Opentelemetry
  - Servidor Grafana

&nbsp;
- **Aplicação em Container com - Docker-compose e Dockerfile**

## Funcionalidades

- **Consulta de Temperatura com retorno Localidade, Celsius, Kelvin e Fahrenheit**
  - O servidor permite consultar a temperatura informando um CEP.
  - Retorno esperado:
```
  {
  "city": "São Paulo",
	"temp_C": 22.3,
	"temp_F": 72.1,
	"temp_K": 295.3
  } 
``` 
  - Sendo city = A cidade do cep informado
  - Sendo temp_F = Fahrenheit
  - Sendo temp_C = Celsius
  - Sendo temp_K = Kelvin  

## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.
   - Foi atulizado a API viaCEP para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
   - Foi utilizado a API WeatherAPI para consultar as temperaturas desejadas: https://www.weatherapi.com/


&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/weather-otel-goexpert.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd weather-otel-goexpert
```
&nbsp;
4. **Rode o docker-compose para buildar e executar toda a stack de observabilidade:**
&nbsp;

```bash 
 docker-compose up
```

&nbsp;


## Como testar localmente:

### Portas
HTTP server on port :8080 <br />

### HTTP
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição.   
  curl --request POST \
  --url http://localhost:8080/ \
  --header 'Content-Type: application/json' \
  --header 'User-Agent: insomnia/10.0.0' \
  --data '{
	"cep": "04911000"
}'

&nbsp;
5. **Acessar o Zipkin para consulta do trace distribuído:**

  - http://localhost:9411/

&nbsp;
6. **Acessar o Jaeger para consulta do trace distribuído:**

  - http://localhost:16686/ 

&nbsp;
7. **Acessar o Grafana para consulta do trace distribuído:**

  - http://localhost:3001/

&nbsp;
8. **Acessar o Prometheus para consulta do trace distribuído:**

  - http://localhost:9090/
  