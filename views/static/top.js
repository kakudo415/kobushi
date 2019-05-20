"use strict";
const ringsListHTML = document.querySelector("#rings > .contents");
const prevPageHTML = document.getElementById("prev-page-button");
const nextPageHTML = document.getElementById("next-page-button");
const newRingFormHTML = document.forms["new-ring-form"];

const sendNewRing = (title, author, desc) => {
  fetch("/", {
    method: "POST",
    body: JSON.stringify({
      title: title,
      author: author,
      description: desc
    })
  }).then((res) => {
    return res.json();
  }).then((json) => {
    location.pathname = `/ring/${json.ring_id}`;
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
  if (ringsListHTML.childElementCount === 25) {
    const nextQuerys = querys;
    nextQuerys["p"] = page + 1;
    nextPageHTML.setAttribute("href", location.pathname + asmQuerys(nextQuerys));
  }
};

init();