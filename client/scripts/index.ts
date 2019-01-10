const websocket = new WebSocket("ws://192.168.0.175:8080/ws");

function isDiv(element: Element): element is HTMLDivElement {
    return element.tagName === "DIV";
}

function isButton(element: Element): element is HTMLButtonElement {
    return element.tagName === "BUTTON";
}

const mapButtonToKeyboard = new Map([
    ["control-up", "up"],
    ["control-left", "left"],
    ["control-right", "right"],
    ["control-down", "down"],
    ["control-a", "a"],
    ["control-b", "b"],
    ["control-c", "c"],
    ["control-d", "d"],
]);

const mapEventToType = new Map([
   ["mousedown", "press"],
   ["mouseup", "release"],
]);

const buttonListeners = ["mousedown", "mouseup"];

function listener(event: any) {
    // function listener(event: MouseEvent) {
    if (event.target === null) {
        return;
    }

    const buttonName = event.target.getAttribute("name");

    const message = {
        key: mapButtonToKeyboard.get(buttonName) || null,
        name: buttonName,
        timestamp: new Date(),
        type: mapEventToType.get(event.type) || null,
    };

    websocket.send(JSON.stringify(message));
    console.log(`Sent: ${JSON.stringify(message)}`);
}

function attachButtonListeners(parent: any): void {
    // function attachButtonListeners(parent: HTMLDivElement): void {
    for (const child of parent.children) {
        if (isDiv(child)) {
            attachButtonListeners(child);
        } else if (isButton(child)) {
            buttonListeners.forEach((eventType) => child.addEventListener(eventType, listener));
        }
        // Fall through
    }
}

window.onload = () => {
    websocket.onopen = (event) => {
        const controller = document.querySelector("#controller");

        attachButtonListeners(controller);
    };
}
