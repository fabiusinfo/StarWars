broker:
	echo "Ejecutando Broker"
	go run broker/Broker.go

ahsoka:
	echo "Ejecutando informante Ahsoka"
	go run informante/Ahsoka.go

almirante:
	echo "Ejecutando informante Almirante Thrawn"
	go run informante/Almirante.go

leia:
	echo "Ejecutando princesa Leia"
	go run leia/Leia.go
	
fulcrum:
	echo "Ejecutando servidor Fulcrum"
	rm -rf servidores/RP/
	mkdir servidores/RP
	go run servidores/Fulcrum.go