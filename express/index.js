import express from "express";
const app = express();

app.use(express.json());

// Text - GET
app.get("/text", (req, res) => {
  res.status(200).send("hello world");
});

// JSON - GET
app.get("/json", (req, res) => {
  res.status(200).send({ hello: "world" });
});

// JSON - POST
app.post("/json", (req, res) => {
  const { hello } = req.body;
  res.status(200).send({ hello });
});

app.listen(3000, () => {
  console.log("Express ready...");
});
