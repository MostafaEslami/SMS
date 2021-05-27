#!/bin/bash

for i in $(eval echo {1..$1})
do 
	 curl -X POST \
	         http://localhost:5740/api/sms/send \
	  -H 'cache-control: no-cache' \
   -H 'content-type: application/x-www-form-urlencoded' \
  -H 'postman-token: 3ea1819d-d2f8-c539-d74f-67a7655ba1c9' \
  -d 'receiver_number=09132957573&code=5487'

  done
