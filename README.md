 # Conversor de Moedas  #
Este projeto é uma API REST para conversão de moedas, escrita em Golang e utilizando o framework Gin e o banco de dados SQLite.

## Rotas ##
A API possui apenas uma rota:
`
GET /exchange/{amount}/{from}/{to}/{rate}
}`


Essa rota converte a quantidade amount da moeda from para a moeda to, utilizando a taxa de câmbio rate. As moedas suportadas são: BRL, USD, EUR e BTC.

Exemplo de uso:
`
GET /exchange/10/BRL/USD/4.50
`

Resposta:  `{
  "valorConvertido": 45,
  "simboloMoeda": "$"
}`


## Instalação e Uso ##
Para rodar a aplicação, é necessário ter o Go instalado na máquina. Além disso, as dependências do projeto devem ser baixadas:
go mod tidy

Para iniciar a aplicação, rode o comando:
`
go run main.go
`

A aplicação será iniciada na porta 8000. Agora, basta fazer requisições para a rota de conversão de moedas.

## Banco de dados ##
O banco de dados utilizado é o SQLite, e as tabelas são criadas automaticamente ao rodar a aplicação. As conversões realizadas são armazenadas no banco de dados para consultas futuras.
