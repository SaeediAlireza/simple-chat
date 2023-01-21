function logout(){
    var xhr = new XMLHttpRequest();
    xhr.open('POST','/api/logout',true);
    xhr.setRequestHeader("Content-Type", "application/json;charset=UTF-8")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4 ) {
            if (xhr.status === 200) {
                window.location.replace("http://localhost:8080/view/login");
            }else{
                window.alert("Not Done!")
            }
        }
    };
    xhr.send()
}
// document.getElementById("chats").addEventListener('onload',getchats);