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
-Se implemento una interfaz para facilitar el uso de comandos, para ingresar un comando se selecciona el numero que especifica la interfaz, en el caso de los informantes, al elegir un comando, se escribe por separado el nombre del planeta, seguido por la ciudad y seguido del valor (no hace falta el valor para el comando DeleteCity)

-Se decidio deshabilitar el funcionamiento de la propagacion de cambios, debido a mal funcionamiento

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
### Instrucciones para el uso del sistema:

Una vez montado todos los procesos:

> Broker -ENTER- para comunicar a los servidores Fulcrum de sus ips

Desde almirante o ahsoka se debiese proceder a añadir ciudades, actualizar nombres, actualizar soldados o eliminar ciudades:

> Informante -1- Addcity 

> Informante -2- UpdateName

> Informante -3- UpdateNumber

> Informante -4- DeleteCity

Despues de haber agregado distintos planetas y ciudades, el proceso Leia, puede solicitar el numero de soldados de una ciudad de un planeta:

> Leia -1- GetNumberRebelds
 
> Leia -Nombre planeta-

> Leia -Nombre ciudad-













