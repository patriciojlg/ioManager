# ioManager

**ioManager** es una herramienta escrita en Go diseñada para facilitar la gestión de archivos en flujos de trabajo automatizados. Permite operaciones como la carga de archivos `.xlsx` desde S3, su transformación en múltiples archivos `.json` y la posterior carga de estos archivos a destinos específicos, como S3.

## 🚀 Características

- **Carga de archivos desde S3**: Descarga archivos `.xlsx` directamente desde Amazon S3.
- **Transformación de datos**: Convierte cada fila de un archivo `.xlsx` en un archivo `.json` independiente.
- **Carga a S3**: Sube los archivos `.json` generados a un bucket de S3 especificado.
- **Procesamiento concurrente**: Utiliza goroutines para manejar múltiples archivos simultáneamente, optimizando el rendimiento.

## 📦 Instalación

1. Clona el repositorio:

   ```bash
   git clone https://github.com/patriciojlg/ioManager.git
   cd ioManager
   ```

2. Compila el proyecto:

   ```bash
   go build -o ioManager
   ```

   Asegúrate de tener Go instalado en tu sistema.

## 🛠️ Uso

### Carga y transformación de archivos `.xlsx` desde S3

```go
files, err := explodeXLSXFromS3("mi-bucket", "ruta/archivo.xlsx")
if err != nil {
    log.Fatal(err)
}
```

### Carga de archivos `.json` a S3

```go
err = uploadJSONFilesConcurrentlyV1(files, "mi-bucket", "ruta/salida", s3Client, 5)
if err != nil {
    log.Fatal(err)
}
```

Asegúrate de configurar correctamente tu cliente de S3 (`s3Client`) antes de utilizar estas funciones.

## 📁 Estructura del Proyecto

- `cmd/ioManager/`: Contiene el punto de entrada principal de la aplicación.
- `controllers/`: Maneja la lógica de negocio y coordina las operaciones.
- `models/`: Define las estructuras de datos utilizadas en el proyecto.
- `providers/`: Implementa la interacción con servicios externos, como S3.
- `settings/`: Contiene la configuración y variables de entorno.
- `shared/utils/`: Proporciona funciones utilitarias comunes.

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si deseas mejorar este proyecto, por favor:

1. Haz un fork del repositorio.
2. Crea una nueva rama para tu funcionalidad o corrección.
3. Realiza tus cambios y haz commit.
4. Envía un pull request describiendo tus cambios.

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.