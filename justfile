run:
    go run ./app

run_replica:
    go run ./app --port 8081 --replicaof localhost 6379

submit:
    git add . && git commit --allow-empty -m 'submit' && git push origin master

test:
    codecrafters test


index:
    http :4221

echo:
    http :4221/echo/grape

notfound:
    http :4221/notfound
