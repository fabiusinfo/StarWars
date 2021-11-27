broker:
	echo "Ejecutando Broker"
	go run middleware/Broker.go

ahsoka:
	echo "Ejecutando informante Ahsoka"
	go run informantes/Ahsoka.go

almirante:
	echo "Ejecutando informante Almirante Thrawn"
	go run informantes/Almirante.go

leia:
	echo "Ejecutando princesa Leia"
	go run princess/Leia.go
	
fulcrum:
	echo "Ejecutando servidor Fulcrum"
	rm -rf servidores/RP/
	mkdir servidores/RP
	go run servidores/Fulcrum.go