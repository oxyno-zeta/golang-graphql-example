#!/bin/sh
#
# To prevent debug code from being accidentally committed, simply add a comment near your
# debug code containing the keyword "not implemented" and this script will abort the commit.
#
if git commit -v --dry-run | grep -E '[nN]ot[- ][iI]mplemented' >/dev/null 2>&1
then
  echo "Trying to commit non-committable code."
  echo "Remove the 'not implemented' string and try again."
  exit 1
else
  exit 0
fi
