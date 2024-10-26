
# test post coffee 
curl -v -X POST  \
-H "Content-Type: application/json" \
-d '{"name" : "test-super-coffee-5" , "varietalId" : 2, "farmerId": 1, "originId": 1 , "processId" :1 , "description" : "something very insightful", "price": 1 , "sca": 8.7 , "acidity" : "normal" , "body" : "strong", "balance" : "medium" , "clarity": "good" , "sweetness" : "intense", "image" :"https://www.antesuncafe.com/static/media/project-img2.8ccdcab64328946462b4.png" }' http://localhost:8080/coffees
