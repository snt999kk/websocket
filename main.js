let socket = new WebSocket("ws://localhost:9191/abc");

socket.onopen = function(e) {
  alert("[open] Соединение установлено");
  alert("Отправляем данные на сервер");
};

let response = ''
socket.onmessage = function(event) {
    response = event.data;
}

let render = document.getElementById('3');
let button = document.getElementById('2');
button.onclick = function()
{
    console.log('got here')
    let val = document.getElementById('1').value;
    if(val != ''){
        socket.send(val)
        render.textContent = 'Message is ' + response;
    }
}