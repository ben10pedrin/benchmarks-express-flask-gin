# Express - Flask - Gin benchmark

Repo from this YouTube video
Todo: Post link

## Results

Todo: Hake a table

Tests were made using apache benchmark tool:

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

## Notes

Tests were made on M1 Macbook Air 8GB RAM
MacOS Ventura 13.0

Node v19.2.0
Python 3.9.6
gunicorn (version 20.1.0)
go1.19.4 darwin/arm64

Server version: Apache/2.4.54 (Unix)
Server built: Sep 30 2022 02:51:01

## Credits & Inspiration

Benawad: https://youtu.be/uFH5S5dUPr8
Fireship: https://youtu.be/N6-Q2dgodLs
Medium: https://medium.com/@jamesjudd_21057/benchmarking-the-request-time-of-laravel-asp-net-core-and-django-7c1c3e9663d
