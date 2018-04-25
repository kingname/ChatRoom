from flask import Flask, jsonify, request, render_template
from flask_cors import CORS
import datetime
import pymongo


handler = pymongo.MongoClient().ChatRoom.Content
app = Flask(__name__)
CORS(app)


@app.route('/')
def index():
    return render_template('index.html')


@app.route('/content')
def content():
    content_obj = handler.find().sort('_id', -1).limit(100)
    resp = []
    for row in content_obj:
        row['_id'] = str(row['_id'])
        resp.append(row)
    return jsonify(resp[::-1])


@app.route('/submit', methods=['POST'])
def submit():
    body = request.get_json()
    name = body.get('username', 'unknown')
    resp = body.get('content', '')
    create_time = body.get('create_time', datetime.datetime.now().strftime('%Y-%m-%d %H:%M:%S'))
    cid = body.get('cid', '')
    handler.insert_one({'username': name, 'content': resp, 'create_time': create_time, 'cid': cid})
    return jsonify({'result': True})
