"use strict";
const kobushisListHTML = document.querySelector("#kobushis > .contents");
const prevPageHTML = document.getElementById("prev-page-button");
const nextPageHTML = document.getElementById("next-page-button");
const newKobushiFormHTML = document.forms["new-kobushi-form"];

const sendNewKobushi = (title, desc) => {
  fetch(location.href, {
    method: "POST",
    body: JSON.stringify({
      title: title,
      description: desc
    })
  }).then((res) => {
    if (!res.ok) {
      throw Error(res.statusText);
    }
    newKobushiFormHTML.elements["title"].value = "";
    newKobushiFormHTML.elements["desc"].value = "";
    return res.json();
  }).then((json) => {
    location.pathname += `/${json.kobushi_id}`;
  }).catch((err) => {
    console.error(err);
  });
};

newKobushiFormHTML.elements["submit"].addEventListener("click", () => {
  const value = (name) => {
    return newKobushiFormHTML.elements[name].value;
  };
  sendNewKobushi(value("title"), value("desc"));
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
  if (kobushisListHTML.childElementCount === 25) {
    const nextQuerys = querys;
    nextQuerys["p"] = page + 1;
    nextPageHTML.setAttribute("href", location.pathname + asmQuerys(nextQuerys));
  }
};

init();