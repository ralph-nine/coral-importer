# coral-importer

## Installation

You can visit the [Releases](https://gitlab.com/coralproject/coral-importer/-/releases) page for the most recent binary release
of the application.

```sh
# Install the coral-importer.
go install .
```

## Strategies

### Legacy

```sh
DATABASE_NAME=coral
MONGO_CONTAINER_ID=mongo-export

# Make the export directory.
mkdir -p "export"

# Dump each collection to the export directory. This operation can take some
# time for larger data sets.
collections=(actions assets comments settings users)
for collection in ${collections[*]}
do
  docker run --rm -ti --link $MONGO_CONTAINER_ID:mongo-export -v $PWD/export:/mnt/export mongo:4.2 mongoexport --host mongo-export -d $DATABASE_NAME -c $collection  -o /mnt/export/${collection}.json
done

# This now provides the export files that can be processed by the importer.
TENANT_ID=c2440817-464e-4a8f-8851-24effd8fee9d
INPUT=data/legacy
OUTPUT=database

coral-importer --quiet legacy --input $INPUT --tenantID $TENANT_ID --output $OUTPUT 2> output.log
```

### Livefyre

```sh
# Perform the import parsing operation.
TENANT_ID=c2440817-464e-4a8f-8851-24effd8fee9d
USER_INPUT=data/livefyre/users.json
COMMENTS_INPUT=data/livefyre/comments.json
OUTPUT=database

coral-importer --quiet livefyre --users $USER_INPUT --comments $COMMENTS_INPUT --tenantID $TENANT_ID --output $OUTPUT 2> output.log
```

## Importing

```sh
# Upload the generated imports to MongoDB.
TARGET_MONGO_CONTAINER=mongo
DATABASE_NAME=coral
CONCURRENCY=$(sysctl -n hw.ncpu)

collections=(commentActions stories users comments)
for collection in ${collections[*]}
do
  if [ ! -f $PWD/$OUTPUT/$collection.json ]
  then
    echo "$PWD/$OUTPUT/$collection.json does not exist, not importing $collection collection"
    continue
  fi

  docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/$collection.json --collection $collection --numInsertionWorkers $CONCURRENCY
done
```

If you're updating documents:

```sh
collections=(commentActions stories users comments)
for collection in ${collections[*]}
do
  if [ ! -f $PWD/$OUTPUT/$collection.json ]
  then
    echo "$PWD/$OUTPUT/$collection.json does not exist, not importing $collection collection"
    continue
  fi

  docker run --rm -ti -v $PWD:/mnt/import --link mongo:$TARGET_MONGO_CONTAINER mongo:4.2 mongoimport --uri=mongodb://mongo/$DATABASE_NAME --file=/mnt/import/$OUTPUT/$collection.json --collection $collection --numInsertionWorkers $CONCURRENCY --mode upsert --upsertFields id
done
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

## Developing

```sh
collections=(actions assets comments settings users)
for collection in ${collections[*]}
do
  head ${collection}.json > ${collection}_sample.json
  head -n1 ${collection}.json > ${collection}_single.json
done
```
