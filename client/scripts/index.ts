const websocket = new WebSocket("ws://192.168.0.175:8080/ws");

function isDiv(element: Element): element is HTMLDivElement {
    return element.tagName === "DIV";
}

function isButton(element: Element): element is HTMLButtonElement {
    return element.tagName === "BUTTON";
}

const mapEventToType = new Map([
   ["mousedown", "press"],
   ["mouseup", "release"],
]);

const buttonListeners = ["mousedown", "mouseup"];

function listener(event: any) {
    if (event.target === null) {
        return;
    }

    const buttonName = event.target.getAttribute("name");

    const message = {
        key: buttonName,
        timestamp: new Date(),
        type: mapEventToType.get(event.type) || null,
    };

    websocket.send(JSON.stringify(message));
}

function attachButtonListeners(parent: any): void {
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
