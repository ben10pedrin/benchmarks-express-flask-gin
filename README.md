# Express - Flask - Gin Benchmark

## Contents

- [Requests per second](#Requests-per-second)
- [Ram usage](#Ram-usage)
- [How tests were made](#How-tests-were-made)
- [Version details](#Version-details)
- [FAQ](#FAQ)

## Results

### Requests per second

![](/images/get-text.png)
![](/images/get-json.png)
![](/images/post-json.png)

### Ram usage

![](/images/standby-ram.png)
![](/images/after-all-tests-ram.png)

## How tests were made

Tests were made using [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) tool:

get-text

```bash
ab -n 100 -c 10 http://0.0.0.0:3000/text > 1.txt
```

get-json

```bash
ab -n 100 -c 10 http://0.0.0.0:3000/json > 1.txt
```

post-json

```bash
ab -p data.json -T application/json -n 100 -c 10 http://0.0.0.0:3000/json > 1.txt
```

To run each server I used the following commands:

### Express

Make sure you have node and npm installed.

```bash
npm install
NODE_ENV=production node index.js
```

### Flask

Create a virtual environment.

```bash
python3 -m venv venv
. venv/bin/activate
pip install -r requirements.txt
```

Run it

```bash
gunicorn -w 1 -b :3000 'index:app'
```

Exit

```bash
deactivate
```

### Gin

Make sure you have golang and gin installed.

```bash
go get github.com/gin-gonic/gin
go build -o index
./index
```

## Version Details

- M1 Macbook Air 8GB RAM
- MacOS Ventura 13.0
- Node v19.2.0
- Python 3.9.6
- gunicorn (version 20.1.0)
- go1.19.4 darwin/arm64
- Server version: Apache/2.4.54 (Unix)
- Server built: Sep 30 2022 02:51:01

## FAQ

### Are you running these on production mode?

Yes. Node is using the `NODE_ENV=production` environment variable. Gin has a `gin.SetMode(gin.ReleaseMode)` which I enabled. I compiled go to a binary using `go build`. Flask is using the [gunicorn server](https://flask.palletsprojects.com/en/2.2.x/deploying/) (not the development one).

### How did you obtained those graphs?

I ran all the tests 5 times each, then saved the output in a `results.xlsx` file (which is in this repository) and Google Sheets to make the graphs.

### Are the apache benchmark results reliable?

Sometimes the tests would give some warnings like `WARNING: The median and mean for the processing time are not within a normal deviation These results are probably not that reliable.`. I don't know exactly why this happened, but those tests were discarded. I only used the ones without the warnings/errors, as you can see in the `results` folder of each benchmark.

### How did you measured ram usage?

I opened the activity monitor app on my mac twice, one time before running tests, and other after the tests.

I closed all other opened apps and left terminal and activity monitor opened during the tests, although I don't think having other apps running on the background would have make such a big difference since servers are < 100mb of ram each one.

### Why don't you use `console.log()`, `print()` or `fmt.Println()` on the servers?

I avoided logging since I think this could affect performance.

### Why does flask's ram usage is split in two processes?

The way gunicorn works is by spawning child processes. As far as I know there's no way around it, so I'm summing the ram usage of the web server plus the ram usage of the worker process. This doesn't happen on express and gin.

### Why didn't you use [mysql](https://github.com/mysqljs/mysql#readme) instead of [mysql2](https://github.com/sidorares/node-mysql2)?

I got [an error](https://stackoverflow.com/questions/50093144/mysql-8-0-client-does-not-support-authentication-protocol-requested-by-server) related to mysql permissions (which seems to be pretty common) in which the solution consists on "risking the securty of the project". So I decided to use mysql2.

### Why don't you use `npm start` to run the node test?

The `npm start` command runs the server as a child process, meaning we have two (`node` and `npm start`) processes, making it more difficult to measure ram usage. Instead we directly call the `node` command.

### Do you think this would work on windows/linux?

Tests were made on macOS.

### Are you running this on HTTP2/HTTPS?

If I'm honest I don't know exactly which http version each server is running. I think that in real life that would be managed by a reverse proxy (nginx?), but correct me if I'm wrong.

### Why don't you run this test on a VM with linux, wouldn't that be more realistic?

Yes. I'm afraid of getting overbilled/IP banned for making a huge amount of requests in a very short period of time (I'd essentialy be DDO'Sing my own IP), so I tested it locally on a mac, plus, I'm a little bit lazy to setup AWS/GCP.

### Why are you using the gunicorn server in flask? Why not other options?

I [think](https://www.google.com/search?q=most+common+python+server+for+flask) gunicorn is a popular option, so apps would be deployed this way in real life. I would want to test [other options](https://flask.palletsprojects.com/en/2.2.x/deploying/#self-hosted-options), maybe in the future. I've heard uWSGI is also very popular.

### Why are you using only 1 worker process with flask?

The [documentation](https://flask.palletsprojects.com/en/2.2.x/deploying/gunicorn/#running) mentions that typically gunicorn runs by spawning various processes. As far as I know node and go are using only one process to run the server, so I thought a fair comparison would be also using a single worker process in flask. Please correct me if I'm wrong, I may do a test with multiple worker processes in the future.

### Why don't you enable [asynchronous events in gunicorn](https://flask.palletsprojects.com/en/2.2.x/deploying/gunicorn/#async-with-gevent-or-eventlet)?

I'm not that familiar with how that works. I [don't think](https://flask.palletsprojects.com/en/2.2.x/deploying/gunicorn/#async-with-gevent-or-eventlet) it would be too beneficial for this particular benchmark, but I may test that in the future.

### Why are you still using flask? Isn't that outdated? You should've chosen [FastAPI](https://fastapi.tiangolo.com/), [Django w/ASGI](https://docs.djangoproject.com/en/4.1/topics/async/) or [Tornado](https://www.tornadoweb.org/en/stable/).

I'm interesting in trying out them all. I just don't have enough time, so I decided to just pick one and I chose flask. I think async servers would be interesting to try, especially when dealing with database connections. The fact that flask didn't do as good on the benchmarks doesn't mean python (the language) is slow, remember we are comparing frameworks.

### Why did you use the [gin](https://github.com/gin-gonic/gin) framework in go?

I've heard it's very popular. I may try other frameworks in the future. Golang has faster frameworks, but I think they're not used in production that much, plus gin is already pretty fast. I'm not using golang default http library because we can't parse url parameters, which I think is crucial when making a real-world API.

### Why don't you include Rust in the benchmarks?

I'd love to do that, but currently I'm not that familiar with Rust.

### Are you biased to any particular technology?

If I'm honest I think node is a great option, in my opinion there are always a ton of packages available on npm, which can be both good and bad, but generally you always find help online if you get stuck in almost anything.

Python is also great, but I've found the node ecosystem has more libraries.

Golang is very fast, but some of it's libraries are not that well mantained, for example [Gorilla Mux](https://github.com/gorilla/mux) was very popular, but has struggled to get mantainers and got archived despite having a ton of stars. Fewer developers know this language, which can be a problem when hiring.

Working with JSON in Golang is more difficult since you have to provide type definitions up-front. In contrast, JavaScript and Python may make it more pleasant to work with JSON, although lack of type-safety can also be a problem.

## Credits & Inspiration

- [Benawad](https://youtu.be/uFH5S5dUPr8)
- [Fireship](https://youtu.be/N6-Q2dgodLs)
- [James Judd](https://medium.com/@jamesjudd_21057/benchmarking-the-request-time-of-laravel-asp-net-core-and-django-7c1c3e9663d)
- [Pete Freitag](https://www.petefreitag.com/item/689.cfm)
- [Hugo Guerrero & Vanessa Ramos](https://developers.redhat.com/blog/2015/04/30/how-to-load-test-and-tune-performance-on-your-api#)
