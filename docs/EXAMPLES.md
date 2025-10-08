# Verification examples

## 1) Get JWT for app-a
./hack/get-token-app-a.sh alice.eng alicepass > token.txt

## 2) Call app-a (allowed)
curl -s -H "Authorization: Bearer $(cat token.txt)" https://gw.example.com/app-a/echo | jq .

## 3) Call app-b (denied if not in group)
curl -i -H "Authorization: Bearer $(cat token.txt)" https://gw.example.com/app-b/echo | head

## 4) Inspect claims
jq -R 'split(".") | .[1] | @base64d | fromjson' < token.txt | jq .
