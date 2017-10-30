set -e
echo "mode: atomic" > coverage.txt

go test -coverprofile=profile.out -covermode=atomic

if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
fi