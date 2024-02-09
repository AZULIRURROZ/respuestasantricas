# «respuestasantricas»

Teniendo go, le mandas como parámetro la dirección de una carpeta o la dirección de muchos archivos .html que sean copias de alguna publicación de los antros, y el programa te muestra como devolución un listado ordenado por cantidad de comentarios de todos los contenidos analizados.

> go run . '/home/azul/Antros/2024-02'

También se puede ejecutar sin parámetros pero después deberán ser ingresados.

Solo está configurado para clones de Rozed o Voxed, pero se puede compatibilizar con más sitios, por ejemplo para que funcione con Temti se le agrega la cadena de texto "tem-comentario" a la variable f_matches.