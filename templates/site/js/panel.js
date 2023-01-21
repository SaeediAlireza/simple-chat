function getchats(){
    
    var xhr = new XMLHttpRequest();
    xhr.open('GET','/api/chats',true);
    
    xhr.onload = function(){
        if(this.status == 200){
            var chats = JSON.parse(this.responseText);
            var chatHtml = ""
            for (var i in chats) {
                
                chatHtml +=
                '<a href="pv/'+chats[i].username+'">'+
                '<div>'+
                '<lu>'+
                '<li>'+ chats[i].name +'</li>'+
                '<li>'+ chats[i].username +'</li>'+
                '</lu>'+
                '</div>'+
                '</a>'
                '<hr>'
                
            }
            
            document.getElementById('chats').innerHTML = chatHtml;
        }
    }
    
    xhr.send()
}