package bd

import (
	"log"

	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es mi objeto de conexión
var MongoCN  = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://moken:951moken951@twitter.fcjvm.mongodb.net/test?retryWrites=true&w=majority");
                                           
// ConectarBD es la funcion que me permite conectar la base de datos
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
 
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	} 
	log.Println("Conexión Exitosa con la BD")
	return client
}

//ChequeoConnection es el pig a la DB
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}