response=($(curl --location --request POST "http://localhost:9001/realms/master/protocol/openid-connect/token" \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'username=admin' \
--data-urlencode 'password=admin' \
--data-urlencode 'client_secret=J9pL1EJGfoEd75p7vt2mOsnuZWwCj9y1' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_id=admin-cli'))

token=`echo ${response[@]} | sed 's/{"access_token":"//g'`;
token=`echo ${token} | sed 's/","expires_in".*//g'`;

# Call api mobile bff
curl --location --request GET "http://localhost:8081/mobile/posts" --header "Content-Type: application/json" --header "Authorization: Bearer ${token}"  | jq

# Call api web bff
curl --location --request GET "http://localhost:3005/users" --header "Content-Type: application/json" --header "Authorization: Bearer ${token}" | jq
