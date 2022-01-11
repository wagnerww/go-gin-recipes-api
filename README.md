## OPEN API(SWAGGER)

Pode baixar o isntalador:
https://github.com/go-swagger/go-swagger/releases

Ou pelo comando no site via docker:
https://goswagger.io/install.html

- Depois só rodar:
    
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
    
  - Rodar o comando de geracão:

        swagger generate spec -o ./swagger.json


## Mongo DB

docker run -d --name mongodb \
    -e MONGO_INITDB_ROOT_USERNAME=admin \
    -e MONGO_INITDB_ROOT_PASSWORD=password \
    -p 27017:27017 \
    -v ${PWD}/home/data:/data/db \
    mongo:4.4.3

### Conectando na base
  mongodb://admin:password@localhost:27017/test