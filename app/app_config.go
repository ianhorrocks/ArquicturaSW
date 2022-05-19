package app

import (
	"os"

	log "github.com/sirupsen/logrus" // 
)

func init() {
	log.SetOutput(os.Stdout) // seteame cual es la salida por la que vas a largar todo lo que te diga
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel) // setea el nivel en el que quiero encontrar la info. Va usar lo que tenga en debug y despues todo lo que tenga por delante
	log.Info("Starting logger system")
}
