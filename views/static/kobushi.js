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

const getQuerys = () => {
  const querys = {};
  if (window.location.search === "") {
    return querys;
  };
  location.search.split("?")[1].split("&").forEach((v) => {
    const query = v.split("=");
    querys[query[0]] = query[1];
  });
  return querys;
};

const asmQuerys = (querys) => {
  if (Object.keys(querys).length === 0) {
    return "";
  }
  const querysArray = [];
  for (let key of Object.keys(querys)) {
    querysArray.push(`${key}=${querys[key]}`);
  }
  return `?${querysArray.join("&")}`;
};

const pageNumber = () => {
  return getQuerys()["p"] || 1;
};

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