## How to create a patch ?

This section is

```sh
git branch tmpsquash PREVIOUSTAG
git checkout tmpsquash
git merge --squash NEWTAG
git commit --no-verify --no-edit
git format-patch master --stdout > patch_PREVIOUSTAG_to_NEWTAG.patch
git checkout master
git branch -D tmpsquash
```
