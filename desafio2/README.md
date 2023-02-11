# Desafio 2 - Multithreading

Como o escopo do desafio era bem aberto, aqui foram contempladas duas soluções diferentes.
Ambas soluções, em seu fundamento, realizam as mesmas tarefas:

- Executam uma consulta de CEP chamando duas APIs distintas ao mesmo tempo.
- A resposta da consulta deve considerar apenas a API que respondeu primeiro, descartando outra chamada.
- A resposta da consulta deve ser exibida na linha de comando, assim com a identificação da API considerada.
- O tempo limite para as duas chamadas é de 1 segundo.

Essas são as aplicações:

- `with-no-api` - Aplicação stand alone para ser rodada no console. Devolve a resposta em uma execução única. A resposta JSON é apenas repassada no mesmo formato de cada API.
- `with-api` - Implementa um serviço de consulta de CEP. Devolve a resposta para o cliente da API e também no console de log do serviço. No log é idenficado qual API que retornou o resultado. A resposta retornada é normalizada para um padrão único de saída.

## Para testar

---
### `with-no-api` 

Executar o comando com o CEP

```
go run ./with-no-api/main.go <cep(apenas números)>
```
Se for executado sem um argumento, um CEP padrão de exemplo será utilizado.

---
### `with-api`

Iniciar o serviço:

```
cd with-api && go run ./cmd/server/main.go
```

Testar no navegador, abrindo a página:

```
localhost:8000/cep/<código do cep (apenas números)>
```
Ou, testar com extensão `REST Client` do VSCode:

- Acessar `./with-api/test/cep.http`.
- Enviar a request (método GET) de teste do arquivo.

Ou, acesse a documentação gerada e teste por lá:

```
http://localhost:8000/docs/index.html
```