# Desafio 1 - Client-Server-API

## Server

- Expõe uma API Rest com um método `get` para consulta do câmbio dólar-real
- Ao receber a chamada para esta API, realiza requisição na API de consulta de câmbio no endereço `https://economia.awesomeapi.com.br/json/last/USD-BRL`
  - Tempo máximo de espera para receber a resposta é de `200 ms` utilizando o package `context`
- Grava o resultado da cotação `bid` do campo JSON em um banco de dados `mysql` com timeout máximo de `10 ms`

## Client

- Faz requisição ao servidor atráves de uma chamada Rest API ao servidor para obter a cotação de câmbio atual dolar-real
  - Timeout é de `300 ms`
- Com o resultado da requisição, escreve em um arquivo `cotacao.txt` o valor da última cotação

## Para testar:
---
### Requisitos:
- docker
- docker-compose
- GO versão 1.19

### Procedimentos:

1 - Primeiro inicie o serviço do banco

``` bash
docker-compose up -d
``` 

Ou, caso o docker compose esteja instalado como um plugin entre os comandos docker

```
docker compose up -d
```

2 - Inicie o servidor:
```
cd server && go run server.go
```

3 - Em outra janela do terminal execute o cliente:
```
go run ./client/client.go
```

### Para encerrar:

- Interrompa o serviço como o comando apropriado no terminal para cada SO (ex.: `ctrl + c`)
- Encerre, a partir do diretório `desafio1`, o serviço do banco com o comando:
```
docker-compose down -v
```
Ou, dependendo do seu ambiente:
```
docker compose down -v
```