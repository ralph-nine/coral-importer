# coral-importer

```sh
go install .
coral-importer livefyre comments --input data/comments_sample.json --output database --tenantID c2440817-464e-4a8f-8851-24effd8fee9d
docker run --rm -ti -v $PWD:/mnt/import --link mongo:mongo mongo:4.2 mongoimport --uri=mongodb://mongo/import --file=/mnt/import/database/comments.json --collection comments
```

If you're updating documents:

```sh
docker run --rm -ti -v $PWD:/mnt/import --link mongo:mongo mongo:4.2 mongoimport --uri=mongodb://mongo/import --file=/mnt/import/database/comments.json --collection comments --mode upsert --upsertFields id
```
