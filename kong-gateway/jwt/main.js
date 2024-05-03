const jwt = require("jsonwebtoken");

const key = "1234";
const secret = "12345";

const token = jwt.sign(secret, key, { expiresIn: "1h" });

console.log(token);
