# How to run

Create a virtual environment

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

# Notes

- Using Guincorn as Web Server
- This won't work on Windows
- Using only 1 process
- It's not async
