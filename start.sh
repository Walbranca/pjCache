#!/usr/bin/env bash
sudo docker-compose up -d --force-recreate pjcache
sudo docker-compose logs -f