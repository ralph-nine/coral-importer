# coral-importer

```sh
go install .
coral-importer livefyre comments --input data/comments_sample.json
docker run --rm -ti -v $PWD:/mnt/import --link mongo:mongo mongo:4.2 mongoimport --uri=mongodb://mongo/import --file=/mnt/import/comments.import.json --collection comments
```

If you're updating documents:

```sh
docker run --rm -ti -v $PWD:/mnt/import --link mongo:mongo mongo:4.2 mongoimport --uri=mongodb://mongo/import --file=/mnt/import/comments.import.json --collection comments --mode upsert --upsertFields id
```
