const express = require("express");
const app = express();

let numberOfRequests = 0;

app.get("/test/:index", (req, res) => {
  numberOfRequests++;
  console.log(
    `index:${req.params.index} - number of requests:${numberOfRequests}`
  );
  return res.json({ oi: "iei" });
});

app.listen(8000);
