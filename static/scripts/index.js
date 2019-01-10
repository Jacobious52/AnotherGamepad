var websocket;

var mode = "play";

function configure(element) {
  if (mode != "configure") {
    return;
  }

  var val = prompt("Enter new value for button:");
  element.setAttribute("value", val);
}

function send(value, state) {
  if (mode != "play") {
    return;
  }

  const message = {
    key: value,
    timestamp: new Date(),
    type: state,
  };
  websocket.send(JSON.stringify(message));
}

function attachButtonListeners() {
  var btns = document.getElementsByClassName("btn");
  for (var i = 0; i < btns.length; i++) {
    let e = btns[i];
    e.addEventListener('touchstart', function () {send(e.getAttribute("value"), "down");});
    // e.addEventListener('touchenter', function () {send(e.getAttribute("value"), "down");});
    // e.addEventListener('mousedown', function () {send(e.getAttribute("value"), "down");});
    e.addEventListener('touchend', function () {send(e.getAttribute("value"), "up");});
    // e.addEventListener('touchleave', function () {send(e.getAttribute("value"), "up");});
    // e.addEventListener('touchcancel', function () {send(e.getAttribute("value"), "up");});
    // e.addEventListener('mouseup', function () {send(e.getAttribute("value"), "up");});

    e.addEventListener('touchstart', function () {configure(e);});
    e.addEventListener('mousedown', function () {configure(e);});
  }

  let config = document.getElementById("config");
  config.addEventListener('change', function(){
    mode = config.checked ? "configure" : "play";
  });
}

window.onload = function () {
  if ("WebSocket" in window) {
    websocket = new WebSocket("ws://" + window.location.host + "/ws");
    websocket.onopen = function () {
      console.log("open");
      attachButtonListeners();
    };
  } else {
    alert("WebSocket NOT supported by your Browser! This is required to run the game.");
  }
}
// 
