# API Mutantes
### Descripción
Se realizó el desarrollo de una API Rest en Golang, la misma expone los siguientes endpoints:

- __Mutant__: Detecta si un individuo es mutante o humano analizando su ADN. 
El ADN debe ser un array de strings, el mismo debe contener 6 items, y cada item debe poseer 6 caracteres. 
Si el individuo es humano devuelve un error 403 Forbidden. 
Si el individuo es mutante devuelve un 200 Ok

   Request:
```   
   POST /mutant
   Body:
   { "dna": ["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"] }
```
   Response:
```
   200 Ok (Es mutante)
   403 Forbidden (Es humano)
   
```

- __Stats__: Devuelve la cantidad de humanos, la cantidad de mutantes validados y además devuelve la proporción entre ellos


   Request:
```   
   GET /stats
```   
   Response:
```   
   200 Ok
   Body:
   {
    "count_mutant_dna": 4,
    "count_human_dna": 3,
    "ratio": 1.333
   }
   
```
___

### Dev
Para poder correr la API de forma local se requiere tener instalado
- Java (requirements: https://docs.oracle.com/javase/7/docs/webnotes/install/windows/windows-system-requirements.html)
- Go (requirements: https://golang.org/doc/install)
- AWS CLI isntalada y configurada
- AWS SDK for GO (instalar mediante comando: go get github.com/aws/aws-sdk-go)
- Gorilla MUX (instalar mediante comando: go get -u github.com/gorilla/mux)

Para poder correr el proyecto se debe parar a raíz del proyecto y ejecutar el comando:
```
go run application.go
```
El proyecto por defecto correra en el puerto :5000. Si este puerto se lo tiene ocupado o no se puede correr la api en el mismo por alguna restricción de su entorno de trabajo, puede modificarlo desde el archivo application.go que se encuentra a raíz en la línea 13.

```
	server := http.ListenAndServe(":5000", router) // Cambiar por el puerto de preferencia
```
___

### Tests

- Todos los test se encuentran a nivel de cada package
- Para correr los tests dentro de cad package correr
``` 
go test
```
___

### Arquitectura
- Se decidió utilizar Amazon AWS como servicio Cloud. 
- Se publicó el servicio de API en ElasticBeanstalk, ya que provee un ambiente seteado para poder correr golang, permitiendome concentrar en la infraestructura. 
Se creó con un auto scaling de 1 a 4 instancias, y con un loadbalancer para distribuir la carga. 
Todas las instancias corren en la región de San Pablo la cuál es la mas cercana y por lo tanto produce menor latencia.
- Para la persistencia se decidió utilizar DynamoDB, un servicio de base de datos NoSQL de AWS, 
con el mismo podemos realizar peticiones de escritura y lectura de alta demanda, 
se configuró un auto scaling para que responda bajo de manda. 
Ademas se configuró un Cluster DAX (DynamoDB Accelerator) con 3 nodos para poder hacer mas rápida la lectura.
![alt text][logo]

[logo]:https://github.com/bgiulianetti/api-mutantes/blob/master/Arquiectura/Architecture.png "Arquitectura"

___
### Endpoints
__/mutants__
``` 
http://apimutantes-env.mmizwaripz.sa-east-1.elasticbeanstalk.com/mutant
``` 
__/stats__
``` 
http://apimutantes-env.mmizwaripz.sa-east-1.elasticbeanstalk.com/stats
``` 
