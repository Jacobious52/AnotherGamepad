var websocket;
var mode = 'play';
const controller = document.getElementById('controller');

var activeButtons = new Map();

function configure(element) {
  if (mode != 'configure') {
    return;
  }

  if (element.classList.contains("inf")) {
    return;
  }
  const val = prompt('Enter new value for button:', element.getAttribute('name'));
  if (val != null) {
    element.setAttribute('name', val);
  }

  const inf = controller.contentDocument.getElementsByClassName("inf");
  for (let i = 0; i < inf.length; i++) {
    const parents = inf[i].getAttribute("data-parent").split(" ");
    const s = parents.map((v) => controller.contentDocument.getElementById(v).getAttribute("name")).join(" ");
    inf[i].setAttribute('name', s);
  }
}

function handleMove(evt) {
  if (mode != 'play') {
    return;
  }
  evt.preventDefault();

  var changes = new Set([]);
  const touches = evt.changedTouches;
  for (let i = 0; i < touches.length; i++) {
    const t = touches[i];
    const el = controller.contentDocument.elementFromPoint(t.clientX, t.clientY);

    if (!el.classList.contains('btn')) {
      continue;
    }

    changes.add(el.id);
  }

  var btns = controller.contentDocument.getElementsByClassName("btn");
  for (let i = 0; i < btns.length; i++) {
    const e = btns[i];
    if (activeButtons.get(e.id) && !changes.has(e.id)) {
      // button up
      send(e, "up");
      activeButtons.set(e.id, false);
    } else if (!activeButtons.get(e.id) && changes.has(e.id)) {
      // button down
      send(e, "down");
      activeButtons.set(e.id, true);
    }
  }
}

function handleEnd(evt) {
  if (mode != 'play') {
    return;
  }
  evt.preventDefault();

  const touches = evt.changedTouches;
  for (let i = 0; i < touches.length; i++) {
    const t = touches[i];
    const el = controller.contentDocument.elementFromPoint(t.clientX, t.clientY);

    if (!el.classList.contains('btn')) {
      continue;
    }
    activeButtons.set(el.id, false);
    send(el, "up");
  }
}

function send(element, state) {
  if (mode != 'play') {
    return;
  }

  const message = {
    key: element.getAttribute('name'),
    timestamp: new Date(),
    type: state,
  };

  websocket.send(JSON.stringify(message));
  if (state == 'down') {
    window.navigator.vibrate(50);
  } else {
    window.navigator.vibrate(0);
  }
}

function attachButtonListeners() {
  controller.contentDocument.addEventListener('touchstart', handleMove, false);
  controller.contentDocument.addEventListener('touchmove', handleMove, false);
  controller.contentDocument.addEventListener('touchend', handleEnd, false);
  controller.contentDocument.addEventListener('touchcancel', handleEnd, false);

  const btns = controller.contentDocument.getElementsByClassName('btn');
  for (let i = 0; i < btns.length; i++) {
    const e = btns[i];
    activeButtons.set(e.id, false);
    e.addEventListener('touchstart', function () {
      configure(e);
    });
  }

  const config = document.getElementById('config');
  config.addEventListener('change', function () {
    mode = config.checked ? 'configure' : 'play';
  });
}

controller.onload = function () {
  if ('WebSocket' in window) {
    websocket = new WebSocket('ws://' + window.location.host + '/ws');
    attachButtonListeners();
  } else {
    alert('WebSocket NOT supported by your Browser! This is required to run the game.');
  }
}
