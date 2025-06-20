qtype() {

    quicktype --package main --debug=print-schema-resolving --top-level TiledMap -s schema "$1" -o ./dist/test.go
}
