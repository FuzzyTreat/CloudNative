#docker build -t pythondemoapp .
#docker run -p 5000:5000 pythondemoapp
# surfa till localhost:5000


from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello():
    return "Yes of course"