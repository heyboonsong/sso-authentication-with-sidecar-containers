token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c";
# Call api mobile bff
mobile=($(curl --location --request GET "http://localhost:8081/mobile/posts" \
--header "Content-Type: application/json" \
--header "Authorization: Bearer ${token}" \
))
echo $mobile

Call api web bff
web=($(curl --location --request GET "http://localhost:8081/web/users" \
--header "Content-Type: application/json" \
--header "Authorization: Bearer ${token}" \
))
echo $web

# Call api desktop bff
desktop=($(curl --location --request GET "http://localhost:8081/desktop" \
--header "Content-Type: application/json" \
--header "Authorization: Bearer ${token}" \
))

echo $desktop
