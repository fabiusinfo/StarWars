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
- Se implementó una interfaz de fácil uso, que consiste en ingresar un número para seleccionar un comando, en el caso de los informantes, al elegir un comando, con su número, se piden por separado sus componentes (planeta, ciudad, valor (en el caso de Delete no pide valor ya que no se requiere)). 
r
Ejemplo:

	Ingrese el número del comando a usar:

	-1- AddCity

	-2- UpdateName
	
	-3- UpdateNumber
	
	-4- DeleteCity
	
	3
	
	Ingrese -Nombre planeta-
	
	earth
	
	Ingrese -Nombre ciudad-
	
	stg
	
	Ingrese -Nuevo valor-
	
	70

- La propagación se hace a los 2 [min] y se muestra un mensaje de una de las máquinas indicando su activación.
- El Broker no tiene  movimiento en cuanto a prints por limpieza, pero es quien redirige la información como debe ser.
- Cada vez que le llega un comando a los Fulcrum se muestra por pantalla el mismo.
- Para los casos en que se intente Actualizar algo de un planeta/ciudad que no exista, los informantes envian los comandos pero no se ejecuta nada particular, no sucede lo mismo con Leia. Cuando se pregunta por un planeta que no existe **se termina el proceso de Leia y del fulcrum al que le pregunta**.
- El merge no está implementado.

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

Una vez montados **todos** los procesos:

> En el Broker presionar -ENTER- para comunicar a los servidores Fulcrum de sus ips

Desde almirante o ahsoka se debiese proceder a añadir ciudades, actualizar nombres, actualizar soldados o eliminar ciudades:

> Informante -1- Addcity 

> Informante -2- UpdateName

> Informante -3- UpdateNumber

> Informante -4- DeleteCity

Después de haber agregado distintos planetas y ciudades, el proceso Leia, puede solicitar el numero de soldados de una ciudad de un planeta:

> Leia -1- GetNumberRebelds
 
> Leia -Nombre planeta-

> Leia -Nombre ciudad-