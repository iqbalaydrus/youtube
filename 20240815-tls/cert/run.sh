openssl genrsa -out ca.key 2048
openssl req -x509 -subj "/CN=Testing Youtube CA" -nodes -key ca.key -days 3650 -out ca.crt
openssl req -newkey rsa:2048 -nodes -subj "/CN=youtube-test.local" -keyout server.key -out server.csr
openssl x509 -req -extfile <(printf "subjectAltName = DNS:youtube-test.local") -in server.csr -out server.crt -CAcreateserial -CA ca.crt -CAkey ca.key -days 3000
#openssl x509 -text -noout -in ca.crt
#openssl x509 -text -noout -in server.crt
