const code = document.getElementById("code")

function add() {
    const url = document.getElementById("url").value
    const password = document.getElementById("code").value
    var person = {
        Url: url,
        Code: password,
    }
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "add", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify(person));
    
    xhr.addEventListener("load", (event) => {
        console.log(xhr.responseText)
        console.log(xhr.status)
        if (xhr.status == 201) {
            location.reload()
        } else {
            document.getElementById("error").innerText = "Failed to add url: " + xhr.statusText;
        }
    });
}

code.addEventListener("keyup", (evnt) => {
    if (evnt.keyCode == 13) {
        add();
    }
});