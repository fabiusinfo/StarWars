# Laboratorio Sistemas Distribuidos
# StarWars

### Integrantes
Fabián Arancibia 201573102-0

Javier Olivares 201373069-8

Katherine Salgado 201610515-8
	
### Desiciones tomadas en el Laboratorio:
- La organización de los procesos es la siguiente:
	- Máquina 1, 10.6.43.41: Broker (8080)
	- Máquina 2, 10.6.43.42: Leia (8080) y servidor Fulcrum 1 (9000)
	- Máquina 3, 10.6.43.43: Ahsoka  (8080) y servidor Fulcrum 2 (9000)
	- Máquina 4, 10.6.43.44: Almirante Thrawn y servidor Fulcrum 3 (9000)
-Se implemento una interfaz para facilitar el uso de comandos


### Instrucciones de ejecución:

dist161:

	> cd StarWars
	
	> make broker
dist162:

	> cd StarWars
	
	> make leia
dist162:

	> cd StarWars
	
	> make fulcrum
dist163:

	> cd StarWars
	
	> make ahsoka
dist163:

	> cd StarWars
	
	> make fulcrum
dist164:

	> cd StarWars
	
	> make almirante
dist164:

	> cd StarWars
	
	> make fulcrum
