#!/bin/sh

if id | grep "^uid=0(" >/dev/null 2>&1; then
  echo Is superuser
else
  echo Is not superuser
fi

IsSuperUser() {
  case $(id) in
  "user=0("*) return 0 ;;
  *) return 1 ;;
  esac
}

IsSuperUser
echo $?
