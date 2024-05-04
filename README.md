# Whois

Programa para recuperar toda la información de un dominio.com.ar

# Instalación

Si tenés golang instado en tu pc entonces sólo necesitas tener instalada la dependencia: whois

## Instalación con golang

```go
go install github.com/k23dev/whoar@latest
```

# Uso


```go
whoar -d [DOMINIO_AR] -o [PATH_DE_SALIDA]
```

Esto guarda un archivo en el path de salida (opcional) con toda la información de realizar un whois.

# Uso con una lista de dominios

Se puede utilizar pasando una lista de dominios
Sólo les hará un whois a aquellos que terminen en .ar

Para que sea más sencillo de comprobar los resultados se crea un archivo especial llamado: **__all_positive_domains_results.txt** que contiene todos los dominios con resultados positivos

**importante:** El archivo debe estar separado por saltos de línea.

El path de salida es opcional.

```go
whoar -f [PATH_DE_ARCHIVO] -o [PATH_DE_SALIDA]
```