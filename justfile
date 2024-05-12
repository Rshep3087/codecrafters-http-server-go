run:
    go run ./app

run_replica:
    go run ./app --port 8081 --replicaof localhost 6379

submit:
    git add . && git commit --allow-empty -m 'submit' && git push origin master

test:
    codecrafters test
