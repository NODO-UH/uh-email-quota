# Plugin Correo Estudiantes

Plugin para brindar informacion acerca de las cuentas de correo de los estudiantes.

## Servicios API

- Consumo de la cuota de correo

## Compilando

Para generar un nuevo binario hay que compilar el proyecto. Para esto se ejecuta el comando:

```bash
make build
```

Con esto tenemos un nuevo binario con el nombre **plugin.bin**

## Ejecucion

El plugin debe estar en el servidor que contiene el servicio de correo, y con acceso a la ejecucion del comando `doveadm`. Ademas debe tener configurada la `API_KEY` como una variable de entorno en el archivo **.env**, el uso del archivo **.env** es obligatorio.

El contenido de la carpeta donde se ejecuta el plugin debe ser el siguiente:

```
plugin_folder
|- .env
|- plugin.bin
|- cert.pem
|- key.unencrypted.pem
```

- **.env** (*requerido*) archivo de configuracion de las variables de entorno
- **plugin.bin** (*requerido*) binario o ejecutable
- **cert.pem** solo es necesario si se va a utilizar el schema HTTPS
- **key.unencrypted.pem** solo es necesario si se va a utilizar el schema HTTPS

Si tenemos todo correcto, ejecutamos el plugin con el comando:

```bash
./plugin.bin --tls-host HOST_IP --tls-port HOST_PORT --tls-key key.unencrypted.pem --tls-certificate cert.pem --scheme=https
```

Donde `HOST_IP` es el IP del servidor y `HOST_PORT` es el puerto de escucha.

Cualquier duda sobre el significado de los parametros, puede ver la eplicacion con el comando:

```bash
./plugin.bin --help
```