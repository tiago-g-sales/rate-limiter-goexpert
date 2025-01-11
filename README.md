# Rate-limiter-goexpert
Sistema API Web com validação quantidade de requisições baseado em parametros, sendo por Ip de origem ou APIKey

# Desafio GOLang rate Limiter - FullCycle 

Aplicação em Go sendo: 
  - Servidor HTTP Rest Client
  - Servidor Jaeger para apresentação do trace distribuído
  - Servidor Prometheus
  - Servidor Opentelemetry
  - Servidor Grafana

&nbsp;
- **Aplicação em Container com - Docker-compose e Dockerfile**

## Funcionalidades

- **Consulta**
  - O servidor permite executar metodo get rota base.

## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.


&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/tiago-g-sales/rate-limiter-goexpert.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd rate-limiter-goexpert
```
&nbsp;
4. **Rode o docker-compose para buildar e executar toda a stack de observabilidade:**
&nbsp;

```bash 
 docker-compose up
```

&nbsp;

## Como testar localmente:

### CONFIGURAÇÕES
  - Variaveis do arquivo docker compose
      - Quantidade de requisições por segundo com origem no mesmo IP  <br />
          HTTP_REQUEST_IP_TPS=2
      - Quantidade de requisições por segundo com a mesma APIKEY      <br />
        Considerar id da APIKEY:Quatidade de requisições por segundo       <br />
          HTTP_REQUEST_APIKEY_TPS=ABC123:3      
      - Tempo de bloqueio após exceder a quantidade de request        <br />
          HTTP_REQUEST_TIME_BLOCK=10s

### Portas
HTTP server on port :8080 <br />

### HTTP COM APIKEY
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição.   
    curl --request GET \
      --url http://localhost:8080/ \
      --header 'API_KEY: ABC123' \
      --header 'User-Agent: insomnia/10.0.0'

### HTTP POR IP DE ORIGEM
 - Execute o curl abaixo ou use um aplicação client REST para realizar a requisição.   
    curl --request GET \
      --url http://localhost:8080/ \
      --header 'User-Agent: insomnia/10.0.0'

&nbsp;

&nbsp;
5. **Acessar o Jaeger para consulta do trace distribuído:**

  - http://localhost:16686/ 

&nbsp;
6. **Acessar o Grafana para consulta do trace distribuído:**

  - http://localhost:3001/

&nbsp;
7. **Acessar o Prometheus para consulta do trace distribuído:**

  - http://localhost:9090/
  