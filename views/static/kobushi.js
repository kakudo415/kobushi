"use strict";
const messagesListHTML = document.querySelector("#messages > .contents");
const prevPageHTML = document.getElementById("prev-page-button");
const nextPageHTML = document.getElementById("next-page-button");
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
    newMessageFormHTML.elements["body"].value = "";
    location.reload();
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

const init = () => {
  const querys = getQuerys();
  const page = Number(pageNumber());
  if (page >= 3) {
    const prevQuerys = querys;
    prevQuerys["p"] = page - 1;
    prevPageHTML.setAttribute("href", location.pathname + asmQuerys(prevQuerys));
  }
  if (page === 2) {
    const prevQuerys = querys;
    delete prevQuerys["p"];
    prevPageHTML.setAttribute("href", location.pathname + asmQuerys(querys));
  }
  if (messagesListHTML.childElementCount === 25) {
    const nextQuerys = querys;
    nextQuerys["p"] = page + 1;
    nextPageHTML.setAttribute("href", location.pathname + asmQuerys(nextQuerys));
  }
};

init();