#!/bin/bash
rsync -av --progress --delete --exclude=.git --exclude=*.test . ~/Copy/github.com/deze333/intvl
