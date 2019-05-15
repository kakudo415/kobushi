"use strict";
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
    location.href = `/ring/${json.ring_id}`;
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