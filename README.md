## OPEN API(SWAGGER)

Pode baixar o isntalador:
https://github.com/go-swagger/go-swagger/releases

Ou pelo comando no site via docker:
https://goswagger.io/install.html

- Depois só rodar:
    
      swagger version

Gerando a doc:

    swagger generate spec -o ./docs/swagger.json

Servindo a doc:

    swagger serve ./docs/swagger.json

Swagger UI:

    swagger serve -F swagger ./docs/swagger.json



### Caso der problema:
  - Executar o vendor:
        
        go mod vendor 
    
  - Rodar o comando de geracão:

        swagger generate spec -o ./docs/swagger.json


### Conectando na base MONGO
  mongodb://admin:password@localhost:27017/test

## EXCUTANDO LOCAL

    PROFILE="dev" go run main.go

## EXECUTANDO TESTES
  