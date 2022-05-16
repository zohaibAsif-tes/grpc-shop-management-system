-Command to run server:
	
	make run-server
	
-Command to run client:
	
	make run-client
	
-A simple grpc service that takes customer's name and list of products as request and generates a bill as a response.

generateBill---------------------------------request-----------------------------------
{
    "bill": {
        "customer":{
            "first_name":"Zohaib",
            "last_name":"Asif"
        },
        "listOfProducts":[
            {"name":"prod1","price":5,"description":"abc"},
            {"name":"prod2","price":5,"description":"abc"}, 
            {"name":"prod3","price":5,"description":"abc"}
        ]            
    }
}

generateBill---------------------------------response----------------------------------
{
    "id": 1,
    "bill": {
        "customer": {
            "first_name": "Zohaib",
            "last_name": "Asif"
        },
        "listOfProducts": [
            {
                "name": "prod1",
                "price": 5,
                "description": "abc"
            },
            {
                "name": "prod2",
                "price": 5,
                "description": "abc"
            },
            {
                "name": "prod3",
                "price": 5,
                "description": "abc"
            }
        ]
    },
    "total": 15,
    "created_at": "2022-05-16 11:13:50.472468798 +0500 PKT"
}
