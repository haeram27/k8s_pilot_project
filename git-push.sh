#!/bin/bash

git commit -m "batch" --all
git push

read -rsp $'\nPress enter to continue...\n'
