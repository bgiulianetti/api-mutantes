const AWS = require('aws-sdk');
const docClient = new AWS.DynamoDB.DocumentClient({region: 'sa-east-1'});
exports.handler = (event, context, callback) =>{
    
    //obtiene info
    var params = {
        TableName : 'individualCount',
        Key:{ "id": "human" }
    };

    var count = 0;
    // Call DynamoDB to read the item from the table
    docClient.get(params, function(err, data) {
      if (err) {
        console.log("Error GET", err);
      } else {
        //var getResponse = JSON.stringify(data.Item);
        console.log("data.Item.count: " + data.Item.count);
        count = parseInt(data.Item.count, 10);
        count++;
        console.log(count);
        
        //Guarda info
        var params = { 
            Item : { 
                id : "human", 
                count : count 
            },  
            TableName : 'individualCount' 
        };
        docClient.put(params, function(err, data){
            if(err) 
                callback(err, null);
            else 
                callback(null, "info guardada");
        })
      }
    });
}