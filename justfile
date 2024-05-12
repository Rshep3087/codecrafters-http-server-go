run:
    go run ./app --directory /Users/ryan/Downloads

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

user-agent:
    http :4221/user-agent User-Agent:Mozilla 

files-exists:
    http :4221/files/Letter.pdf

files-not-exists: 
    http :4221/files/does-not-exist
