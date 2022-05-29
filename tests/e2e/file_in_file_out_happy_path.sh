
set -e

go run main.go -in-file tests/e2e/fixtures/filein.txt -out-file /tmp/usynctest.txt
diff /tmp/usynctest.txt tests/e2e/fixtures/file_in_file_out_happy_path.expected