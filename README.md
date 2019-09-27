# coral-importer

```sh
# Install the coral-importer.
go install .

# Perform the import parsing operation.
TENANT_ID=c2440817-464e-4a8f-8851-24effd8fee9d
USER_INPUT=data/comments_sample.json
COMMENTS_INPUT=data/comments_sample.json
OUTPUT=database

coral-importer --quiet livefyre comments --users $USER_INPUT --comments $COMMENTS_INPUT --tenantID $TENANT_ID --output $OUTPUT 2> output.log

# Upload the generated imports to MongoDB.
TARGET_MONGO_CONTAINER=mongo
DATABASE_NAME=coral
CONCURRENCY=$(sysctl -n hw.ncpu)

docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/comments.json --collection comments --numInsertionWorkers $CONCURRENCY
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/stories.json  --collection stories --numInsertionWorkers $CONCURRENCY
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/users.json  --collection users --numInsertionWorkers $CONCURRENCY
```

If you're updating documents:

```sh
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/comments.json --collection comments --numInsertionWorkers $CONCURRENCY --mode upsert --upsertFields id
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/stories.json  --collection stories --numInsertionWorkers $CONCURRENCY --mode upsert --upsertFields id
docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/users.json  --collection users --numInsertionWorkers $CONCURRENCY --mode upsert --upsertFields id
```

## Tricks

Print all the active operations on your database with messages.

```js
db.currentOp()
  .inprog.filter(op => Boolean(op.msg))
  .map(op => ({ ns: op.ns, msg: op.msg, command: op.command }));
```

## Benchmarks

|                    |       |
| ------------------ | ----- |
| Running importer   | 1m30s |
| Importing Comments | 4m30s |
| Importing Stories  | 4s    |
| Importing Users    | 30s   |
| Rebuilding Indexes | 20m   |
