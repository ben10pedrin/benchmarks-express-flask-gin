# How to run

Make sure you have node and npm installed.

```bash
npm install
NODE_ENV=production node index.js
```

# Note

It's important to avoid using `npm start` for the tests, as this runs the server as a child process, making it more difficult to measure ram usage.
