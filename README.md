# Observation Zones Importer

## Development

```bash
make install_linters
make lint
```

## Build

The following command will create executable files in `./build` directory:

```bash
make build
```

## Run

```bash
LOGIN="login" PASSWORD="password" OUTPUT_DIR=./public ./build/observation-zones-importer-mac
```

Note: if you choose Linux then you need to use `./build/observation-zones-importer-linux` instead.

## Test

```bash
make test
```
