document.getElementById('text').focus()
loginform = document.getElementById("login");
loginform.addEventListener('submit',login,false)

function login(){
    debugger
    var xhr = new XMLHttpRequest();
    xhr.open('POST','/api/msg',true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 ) {
            if (xhr.status === 200) {
                document.getElementById('text').focus()
               
            }else{
                window.alert("Usernme or Password is wrong try again...")
            }
        }
    };
    xhr.send(JSON.stringify({   "text": document.getElementById('text').value,
                                "senderuname": document.getElementById('sendername').value,
                                "receveruname": document.getElementById('recevername').value }))
    
}



