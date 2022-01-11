## OPEN API(SWAGGER)

Pode baixar o isntalador:
https://github.com/go-swagger/go-swagger/releases

Ou pelo comando no site via docker:
https://goswagger.io/install.html

Depois só rodar:
    swagger version

Gerando a doc:

  swagger generate spec -o ./swagger.json

Servindo a doc:

    swagger serve ./swagger.json

Swagger UI:

    swagger serve -F swagger ./swagger.json



### Caso der problema:
  - Executar o vendor:
    go mod vendor 
    
  Rodar o comando de gereácão
    swagger generate spec -o ./swagger.json
