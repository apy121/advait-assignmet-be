#!/bin/bash

echo "Testing Sign Up"
curl -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password"}' http://localhost:8080/signup

echo "Testing Sign In"
TOKEN=$(curl -s -X POST -H "Content-Type: application/json" -d '{"email":"test@example.com", "password":"password"}' http://localhost:8080/signin | jq -r '.token')

echo "Testing Protected Endpoint"
curl -X GET -H "Authorization: Bearer $TOKEN" http://localhost:8080/protected

echo "Testing Token Revocation"
curl -X POST -H "Authorization: Bearer $TOKEN" http://localhost:8080/revoke

echo "Testing Token Refresh"
curl -X POST -H "Authorization: Bearer $TOKEN" http://localhost:8080/refresh