var websocket;
var mode = "play";
const controller = document.getElementById("controller");

function configure(element) {
  if (mode != "configure") {
    return;
  }

  const val = prompt("Enter new value for button:", element.getAttribute("name"));
  element.setAttribute("name", val);
}

function send(element, state) {
  if (mode != "play") {
    return;
  }

  const message = {
    key: element.getAttribute("name"),
    timestamp: new Date(),
    type: state,
  };


  websocket.send(JSON.stringify(message));
  if (state == "down") {
    window.navigator.vibrate(50);
  } else {
    window.navigator.vibrate(0);
  }
}

function attachButtonListeners() {
  const btns = controller.contentDocument.getElementsByClassName("btn");
  for (let i = 0; i < btns.length; i++) {
    const e = btns[i];
    e.addEventListener('touchstart', function () {send(e, "down");}, false);
    e.addEventListener('touchend', function () {send(e, "up");}, false);
    e.addEventListener('touchcancel', function () {send(e, "up");}, false);
    e.addEventListener('touchstart click', function () {configure(e);}, false);
  }

  const config = document.getElementById("config");
  config.addEventListener('change', function () {
    mode = config.checked ? "configure" : "play";
  });
}

controller.onload = function () {
  if ("WebSocket" in window) {
    Notification.requestPermission();
    websocket = new WebSocket("ws://" + window.location.host + "/ws");
    attachButtonListeners();
  } else {
    alert("WebSocket NOT supported by your Browser! This is required to run the game.");
  }
}
