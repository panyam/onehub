import os
from ipdb import set_trace
from datetime import datetime
from collections import defaultdict
import itertools
import typesense as ts
import time
import random
import csv
import requests

tsclient = ts.Client({"api_key": "my_api_key", "nodes": [{"host": "localhost", "port": 8108, "protocol": "http"}]})

def tsreq(path, method="GET", body=None):
    methfunc = getattr(requests, method.lower().strip())
    headers = { "X-TYPESENSE-API-KEY": "my_api_key" }
    json=False
    if body is None or type(body) is dict:
        json=True
        headers["Content-Type"] = "application/json"
    url = f"http://localhost:8108/{path}"
    if json:
        return methfunc(url, headers=headers, json=body)
    else:
        return methfunc(url, headers=headers, data=body)

def sendmsg(users, tid, msg):
    creator = random.choice(users)["id"]
    auth = f"{creator}:{creator}123"
    payload = {
            "messages": [{
                "topic_id": tid,
                "base": {
                    "creator_id": creator
                },
                "content_base": {
                    "content_type": "text/plain",
                    "content_text": msg,
                },
            } for msg in msgs]}
    resp = requests.post(f"http://{auth}@localhost:7080/api/v1/topics/{tid}/messages", json= payload)
    rj = resp.json()
    return rj["messages"]

MSG_START_TIME = time.time() - (3600*24*7)

def sendmsgs(users, tid, msgs):
    auth = f"admin:admin123"
    payload = {
            "allow_userids": True,
            "messages": [{
                "topic_id": tid,
                "base": {
                    "creator_id": random.choice(users)["id"],
                },
                "content_base": {
                    "content_type": "text/plain",
                    "content_text": msg,
                },
            } for msg in msgs]}
    resp = requests.post(f"http://{auth}@localhost:7080/api/v1/topics/{tid}/messages", json= payload)
    rj = resp.json()
    return rj["messages"]

def ensure_users(nusers=100):
    out = []
    auth = f"admin:admin123"
    for i in range(nusers):
        userid = f"ltuser{i}"
        resp = requests.get(f"http://{auth}@localhost:7080/api/v1/users/{userid}")
        if resp.status_code == 200:
            user = resp.json()["user"]
            out.append(user)
            print("Found User: ", user)
        else:
            # create it
            user = { "id": userid, "name": make_random_name() }
            resp = requests.post(f"http://{auth}@localhost:7080/api/v1/users", json = {"user": user})
            user = resp.json()["user"]
            out.append(user)
            print("Created User: ", user)
    return out

def extract_topic_title(msg, max_length=40):
    parts = [m.strip() for m in msg.replace("\t", " ").split(" ") if m.strip()]
    if len(parts) == 1:
        return parts[0][:max_length]
    else:
        out = []
        for i,part in enumerate(parts):
            out.append(part)
            if len(" ".join(out)) > max_length:
                out.pop()
                break
        return " ".join(out)

def ensure_topics(users, ntopics=100):
    lines = load_chat_messages_dataset()
    topics = {}
    currtid = None
    for tid, msg, _ in lines:
        if len(topics) >= ntopics: break
        creator = random.choice(users)["id"]
        auth = f"{creator}:{creator}123"
        if tid != currtid:
            currtid = tid
            # create new topic
            tid = f"lt{tid}"
            topicname = extract_topic_title(msg)
            topic = {"topic": { "id": tid, "name": topicname, }}
            resp = requests.get(f"http://{auth}@localhost:7080/api/v1/topics/{tid}")
            if resp.status_code == 200:
                topic = resp.json()["topic"]
                print("Found Topic: ", topic)
            else:
                resp = requests.post(f"http://{auth}@localhost:7080/api/v1/topics", json = topic)
                topic = resp.json()["topic"]
                print("Created Topic: ", topic)
            topics[topic["base"]["id"]] = topic
    return topics

def grouped_messages():
    # should use groupby but cant get it to work
    grouped = defaultdict(list)
    lines = load_chat_messages_dataset()
    for tid, msg in lines:
        grouped[tid].append(msg)
    grouped = list(grouped.items())
    grouped.sort()
    return grouped

def to_datetime(t):
    # tmfmt = "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
    dt = datetime.fromtimestamp(t)
    return dt.isoformat() + "Z"

def create_random_messages(users, tid, msgs):
    time_between_messages = random.random() * 300
    num_users = 10 + int(random.random() * 90)
    topic_users = random.sample(users, num_users)
    start_time = MSG_START_TIME + (random.random() * 3600)
    messages = [{
        "topic_id": tid,
        "base": {
            "creator_id": random.choice(topic_users)["id"],
            "created_at": to_datetime(start_time + (time_between_messages * i)),
            "updated_at": to_datetime(start_time + (time_between_messages * i)),
        },
        "content_base": {
            "content_type": "text/plain",
            "content_text": msg,
        },
    } for i,msg in enumerate(msgs)]
    return messages

def importmsgs(msgs):
    print(f"Importing {len(msgs)} Messages")
    auth = f"admin:admin123"
    payload = { "messages": msgs }
    resp = requests.post(f"http://{auth}@localhost:7080/api/v1/messages:import", json= payload)
    rj = resp.json()
    return rj["messages"]

def generate_batch_messages(users, topics):
    grouped = grouped_messages() 
    starttime = time.time()
    count = 0
    batch = []
    MAX_BATCH_SIZE = 5000
    for tid, msgs in grouped:
        print(f"Generating {len(msgs)} messages for topic: {tid}")
        # find a random time between messages ssay within 5 minutes
        tid = 1 + (int(tid) % len(topics))
        topic_id = f"lt{tid}"
        messages = create_random_messages(users, topic_id, msgs)
        batch.extend(messages)
        if len(batch) > MAX_BATCH_SIZE:
            importmsgs(batch)
            batch = []
        # sendmsgs(users, topic_id, msgs)
        count += len(msgs)
    if batch:
        importmsgs(batch)
        batch = []
    endtime = time.time()
    print(f"Generated {count} messages in {endtime - starttime} seconds")

def load_chat_messages_dataset():
    # Obtained from https://www.kaggle.com/datasets/arnavsharmaas/chatbot-dataset-topical-chat
    chatset_url = "https://www.kaggle.com/datasets/arnavsharmaas/chatbot-dataset-topical-chat"
    if not os.path.isfile("./configs/chatmessages.csv"):
        raise Exception(f"Download the dataset from {chatset_url}")
    return list(csv.reader(open("./configs/chatmessages.csv")))

def generate_messages(users, topics, start=0, count=1000):
    lines = load_chat_messages_dataset()
    starttime = time.time()
    lasttid = None
    for tid, msg in lines[start:count]:
        tid = 1 + (int(tid) % len(topics))
        if tid is not lasttid:
            print(f"Generating messages for topic: {tid}")
        sendmsg(users, f"lt{tid}", msg)
        lasttid = tid
    endtime = time.time()
    print(f"Generated {count} messages in {endtime - starttime} seconds")

def make_random_name():
  adj = random.choice(ADJECTIVES)
  animal = random.choice(ANIMALS)
  adj = adj[0].upper() + adj[1:]
  animal = animal[0].upper() + animal[1:]
  return f"{adj} {animal}"

ADJECTIVES = [
  "adaptable",
  "adventurous",
  "affable",
  "affectionate",
  "agreeable",
  "ambitious",
  "amiable",
  "amicable",
  "amusing",
  "brave",
  "bright",
  "broadminded",
  "calm",
  "careful",
  "charming",
  "communicative",
  "compassionate	",
  "conscientious",
  "considerate",
  "convivial",
  "courageous",
  "courteous",
  "creative",
  "decisive",
  "determined",
  "diligent",
  "diplomatic",
  "discreet",
  "dynamic",
  "easygoing",
  "emotional",
  "energetic",
  "enthusiastic",
  "exuberant",
  "fair-minded",
  "faithful",
  "fearless",
  "forceful",
  "frank",
  "friendly",
  "funny",
  "generous",
  "gentle",
  "good",
  "gregarious",
  "hard-working",
  "helpful",
  "honest",
  "humorous",
  "imaginative",
  "impartial",
  "independent",
  "intellectual",
  "intelligent",
  "intuitive",
  "inventive",
  "kind",
  "loving",
  "loyal",
  "modest",
  "neat",
  "nice",
  "optimistic",
  "passionate",
  "patient",
  "persistent	",
  "pioneering",
  "philosophical",
  "placid",
  "plucky",
  "polite",
  "powerful",
  "practical",
  "pro-active",
  "quick-witted",
  "quiet",
  "rational",
  "reliable",
  "reserved",
  "resourceful",
  "romantic",
  "self-confident",
  "self-disciplined",
  "sensible",
  "sensitive",
  "shy",
  "sincere",
  "sociable",
  "straightforward",
  "sympathetic",
  "thoughtful",
  "tidy",
  "tough",
  "unassuming",
  "understanding",
  "versatile",
  "warmhearted",
  "willing",
  "witty",
]

ANIMALS = [
  'Tiger',
  'Lion',
  'Elephant',
  'Leopard',
  'Panther',
  'Cheetah',
  'Wolf',
  'Jaguar',
  'Hyena',
  'Giraffe',
  'Deer',
  'Zebra',
  'Gorilla',
  'Monkey',
  'Chimpanzee',
  'Bear',
  'Wild Boar',
  'Hippopotamus',
  'Kangaroo',
  'Rhinoceros',
  'Crocodile',
  'Panda',
  'Squirrel',
  'Mongoose',
  'Porcupine',
  'Koala Bear',
  'Wombat',
  'Meerkat',
  'Otter',
  'Hedgehog',
  'Possum',
  'Chipmunk',
  'Squirrel',
  'Raccoon',
  'Jackal',
  'Hare',
  'Mole',
  'Rabbit',
  'Alligator',
  'Monitor Lizard',
  'Oryx',
  'Elk',
  'Badger',
  'Dinosaur',
  'Pangolin',
  'Mole',
  'Okapi',
  'Camel',
  'Wild cat',
  'Coyote',
  'Bison',
  'African Elephant',
  'Aardvark',
  'Antelope',
  'Alpine Goat',
  'Komodo Dragon',
  'Bearded Dragon',
  'Royal Bengal Tiger',
  'Flying Squirrel',
  'Emu',
  'Eagle',
  'Eel',
  'Asiatic Lion',
  'Armadillo',
  'Beaver',
  'Emperor Penguin',
  'Baboon',
  'Bat',
  'Chameleon',
  'Bull',
  'Giant Panda',
  'Chihuahua',
  'Orangutan',
  'Chinchillas',
  'Hawk',
  'Iguana',
  'Ibis',
  'Ibex',
  'King Cobra',
  'Jellyfish',
  'Goose',
  'Walrus',
  'Seal',
  'Skink',
  'Markhor',
  'Falcon',
  'Bull Shark',
  'Arctic Wolf',
  'Owl',
  'Bulbul',
  'Bobcat',
  'Guinea Pig',
  'Yak',
  'Reindeer',
  'Moose',
  'Puma',
  'Okapi',
  'Marten',
  'Squirrel Monkey',
  'Caracal'
]

