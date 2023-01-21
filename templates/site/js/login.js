loginform = document.getElementById("login");
loginform.addEventListener('submit',login,false)

function login(){
    debugger
    var xhr = new XMLHttpRequest();
    xhr.open('POST','/api/login',true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 ) {
            if (xhr.status === 200) {
                window.location.replace("/view/panel");
            }else{
                window.alert("Usernme or Password is wrong try again...")
            }
        }
    };
    xhr.send(JSON.stringify({ "username": document.getElementById('username').value,
                                "pass": document.getElementById('password').value }))
    
}
