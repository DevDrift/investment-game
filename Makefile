GO=${GOROOT}/bin/go

git-update:
	git rm -rf --cached .
	git add .
