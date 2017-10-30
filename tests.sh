set -e
echo "" > coverage.txt

go test -coverprofile=profile.out -covermode=atomic
