# Manage git


## remove/replace a tag
```shell
git push origin :refs/tags/<tagname>
Replace the tag to reference the most recent commit
git tag -fa <tagname>
Push the tag to the remote origin
git push origin --tags
#check
git log --oneline
```

