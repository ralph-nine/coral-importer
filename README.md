# coral-importer

```sh
TARGET_MONGO_CONTAINER=mongo
TENANT_ID=c2440817-464e-4a8f-8851-24effd8fee9d
INPUT=data/comments_sample.json
OUTPUT=database
DATABASE_NAME=coral

# Install the coral-importer.
go install .

# Perform the import parsing operation.
coral-importer livefyre comments --input $INPUT --output $OUTPUT --tenantID $TENANT_ID

# Upload the generated imports to MongoDB.
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/comments.json --collection comments
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/stories.json  --collection stories
```

If you're updating documents:

```sh
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/comments.json --collection comments --mode upsert --upsertFields id
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/stories.json  --collection stories  --mode upsert --upsertFields id
```
