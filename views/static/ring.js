"use strict";
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