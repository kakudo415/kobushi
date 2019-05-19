"use strict";
const newMessageFormHTML = document.forms["new-message-form"];

const sendNewMessage = (body) => {
  fetch(location.href, {
    method: "POST",
     body: JSON.stringify({
       body: body
     })
  }).then((res) => {
    if (!res.ok) {
      throw Error(res.statusText);
    }
  }).catch((err) => {
    console.error(err);
  });
};

newMessageFormHTML.elements["submit"].addEventListener("click", () => {
  const value = (name) => {
    return newMessageFormHTML.elements[name].value;
  };
  sendNewMessage(value("body"));
});