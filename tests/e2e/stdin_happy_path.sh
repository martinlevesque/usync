set -e

echo "helloworld" | go run main.go -stdin=true -out-file /tmp/usynctest.txt
diff /tmp/usynctest.txt tests/e2e/fixtures/stdin_happy_path.expected