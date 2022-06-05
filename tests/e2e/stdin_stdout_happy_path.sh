set -e

echo -n "helloworld" | go run main.go -stdin=true -stdout=true > /tmp/usynctest.txt
diff /tmp/usynctest.txt tests/e2e/fixtures/stdin_stdout_happy_path.expected