"use strict";
const ringsListHTML = document.querySelector("#rings > .contents");
const prevPageHTML = document.getElementById("prev-page-button");
const nextPageHTML = document.getElementById("next-page-button");
const newRingFormHTML = document.forms["new-ring-form"];

const sendNewRing = (title, author, desc) => {
  fetch(location.href, {
    method: "POST",
    body: JSON.stringify({
      title: title,
      author: author,
      description: desc
    })
  }).then((res) => {
    if (!res.ok) {
      throw Error(res.statusText);
    }
    newRingFormHTML.elements["title"].value = "";
    newRingFormHTML.elements["author"].value = "";
    newRingFormHTML.elements["desc"].value = "";
    return res.json();
  }).then((json) => {
    location.pathname += `/ring/${json.ring_id}`;
  }).catch((err) => {
    console.error(err);
  });
};

newRingFormHTML.elements["submit"].addEventListener("click", () => {
  const value = (name) => {
    return newRingFormHTML.elements[name].value;
  };
  sendNewRing(value("title"), value("author"), value("desc"));
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
  if (ringsListHTML.childElementCount === 25) {
    const nextQuerys = querys;
    nextQuerys["p"] = page + 1;
    nextPageHTML.setAttribute("href", location.pathname + asmQuerys(nextQuerys));
  }
};

init();