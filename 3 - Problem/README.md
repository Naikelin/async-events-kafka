# El gremio de Sopaipilleros

El gremio de Sopaipilleros establece un mecanismo en donde todos los que quieran vender Sopaipillas deberán hacerlo de forma legal. Aquél que sea descubierto vendiendo Sopaipillas de forma ilegal, será condenado a carcel con una sentencia mínima de 100 años.

Para unirse al gremio de Sopaipilleros, cualquiera puede hacerlo, de forma libre. Sin embargo, para hacerlo, el interesado se debe conectar a un servidor central. Este servidor se encargará de recibir la ubicación de los carritos, junto a la cantidad de ventas de sopaipillas por parte de cada uno.

A hacer:
    - Kafka -> 2 tópicos 
        - Ubicación + Ventas
    - Servidor
        - Recibe ubicación y ventas
    - Se debe procesar a las 5 de la tarde las ventas y sacar estadísticas
    - La ubicación se debe procesar "casi" en tiempo real

Existe un grupo de Sopaipilleros que quiere controlar el gremio de forma "legal". Para eso contrataron a un "Hacker" que ha logrado tener acceso al Broker de Kafka. Es así como han decidido mantener "territorios" en donde no pueden existir 

