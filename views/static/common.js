"use strict";

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